package resolver

import (
	"context"
	"errors"
	"fmt"
	"musicboxd/api/middleware"
	"musicboxd/database"
	"musicboxd/graph/model"
	"musicboxd/hlp"
	"slices"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate go run github.com/99designs/gqlgen generate
// This file will not be regenerated automatically.
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

type AverageRating struct {
	AverageRating float64 `bson:"averageValue"`
	Count         int     `bson:"count"`
}

type UserReviewNumbers struct {
	AlbumReviews int64 `json:"albumReviews"`
	TrackReviews int64 `json:"trackComments"`
}

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(middleware.GinContextKey)
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}

func ValidateJWT(ctx context.Context) (*hlp.CustomClaims, error) {
	ginCtx, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	jwt := ginCtx.GetHeader("Authorization")
	if jwt == "" {
		return nil, err
	}

	cc, err := hlp.ValidateJWT(jwt)
	if err != nil {
		return nil, err
	}

	return &cc, nil
}

func GetUserAccessToken(ctx context.Context) (string, error) {
	// TODO
	// cache access tokens for a certain amount of time, based on the username
	// also add limit to number of cached tokens - should not be a problem for now tho
	cc, err := ValidateJWT(ctx)
	if err != nil {
		return "", err
	}

	coll := database.GetDB().GetCollection("users")
	user := coll.FindOne(ctx, bson.M{"_id": cc.UserID})
	fmt.Println(user)
	if user.Err() != nil {
		return "", user.Err()
	}

	res := model.User{}
	err = user.Decode(&res)
	if err != nil {
		return "", err
	}

	return *res.Tokens.AccessToken, nil
}

func GetPreloads(ctx context.Context) []string {
	return GetNestedPreloads(
		graphql.GetOperationContext(ctx),
		graphql.CollectFieldsCtx(ctx, nil),
		"",
	)
}

func GetNestedPreloads(ctx *graphql.OperationContext, fields []graphql.CollectedField, prefix string) (preloads []string) {
	for _, column := range fields {
		prefixColumn := GetPreloadString(prefix, column.Name)
		preloads = append(preloads, prefixColumn)
		preloads = append(preloads, GetNestedPreloads(ctx, graphql.CollectFields(ctx, column.Selections, nil), prefixColumn)...)
	}
	return
}

func GetPreloadString(prefix, name string) string {
	if len(prefix) > 0 {
		return prefix + "." + name
	}
	return name
}

func isFieldRequested(ctx context.Context, fieldName string) bool {
	fields := GetPreloads(ctx)
	return slices.Contains(fields, fieldName)
}

func AddLikeDislike(ctx context.Context, userID primitive.ObjectID, itemID string, action string, collection string) (*mongo.SingleResult, error) {
	if action != "like" && action != "dislike" {
		return nil, errors.New("invalid action")
	}

	action += "s"
	oppositeAction := "likes"
	if action == "likes" {
		oppositeAction = "dislikes"
	}

	convertedItemID, err := primitive.ObjectIDFromHex(itemID)
	if err != nil {
		return nil, err
	}

	var projection *bson.M
	if collection == "reviews" {
		projection = database.GetReviewProjection(userID)
	}
	if collection == "comments" {
		projection = database.GetCommentProjection(userID)
	}
	if projection == nil {
		return nil, errors.New("invalid collection")
	}

	coll := database.GetDB().GetCollection(collection)
	result := coll.FindOneAndUpdate(
		ctx,
		bson.M{"_id": convertedItemID},
		bson.M{
			"$addToSet": bson.M{
				action: userID,
			},
			"$pull": bson.M{
				oppositeAction: userID,
			},
		},
		options.FindOneAndUpdate().
			SetReturnDocument(options.After).
			SetProjection(projection),
	)
	if result.Err() != nil {
		return nil, result.Err()
	}

	return result, nil
}

func GetAverageRaiting(ctx context.Context, itemID string) (*AverageRating, error) {
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"itemId": itemID,
			},
		},
		{
			"$group": bson.M{
				"_id": nil,
				"averageValue": bson.M{
					"$avg": "$value",
				},
				"count": bson.M{
					"$sum": 1,
				},
			},
		},
	}

	coll := database.GetDB().GetCollection("reviews")
	cursor, err := coll.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	result := AverageRating{
		AverageRating: 0,
		Count:         0,
	}

	if !cursor.Next(context.Background()) {
		return &result, nil
	}

	if err := cursor.Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode result: %w", err)
	}

	return &result, nil
}

func FillCommentUsers(ctx context.Context, comments []*model.Comment) error {
	if len(comments) == 0 {
		return nil
	}

	// Unique user IDs
	userIDMap := make(map[string]bool)
	for _, comment := range comments {
		if comment.UserID != nil && *comment.UserID != "" {
			userIDMap[*comment.UserID] = true
		}
	}

	userIDs := make([]primitive.ObjectID, 0, len(userIDMap))
	for id := range userIDMap {
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return err
		}
		userIDs = append(userIDs, objID)
	}

	coll := database.GetDB().GetCollection("users")
	usersCursor, err := coll.Find(
		ctx,
		bson.M{"_id": bson.M{"$in": userIDs}},
	)
	if err != nil {
		return err
	}
	defer usersCursor.Close(ctx)

	var users []*model.UserResponse
	if err = usersCursor.All(ctx, &users); err != nil {
		return err
	}

	userMap := make(map[string]*model.UserResponse)
	for _, user := range users {
		userMap[user.ID] = user
	}

	for _, comment := range comments {
		if user, ok := userMap[*comment.UserID]; ok {
			comment.User = user
		}
	}

	return nil
}

func GetUserReviewNumbers(ctx context.Context, userID primitive.ObjectID) (*UserReviewNumbers, error) {
	coll := database.GetDB().GetCollection("reviews")
	counts := UserReviewNumbers{
		AlbumReviews: 0,
		TrackReviews: 0,
	}

	albumCount, err := coll.CountDocuments(ctx, bson.M{"userId": userID, "itemType": "album"})
	if err != nil {
		return nil, err
	}
	counts.AlbumReviews = albumCount

	trackCount, err := coll.CountDocuments(ctx, bson.M{"userId": userID, "itemType": "track"})
	if err != nil {
		return nil, err
	}
	counts.TrackReviews = trackCount

	return &counts, nil
}
