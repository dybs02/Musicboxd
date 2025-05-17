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

func GetReviewProjection(userID primitive.ObjectID) *bson.M {
	return &bson.M{
		// Calculate fields
		"likesCount":    bson.M{"$size": bson.M{"$ifNull": bson.A{"$likes", bson.A{}}}},
		"dislikesCount": bson.M{"$size": bson.M{"$ifNull": bson.A{"$dislikes", bson.A{}}}},
		"userReaction": bson.M{
			"$cond": bson.A{
				bson.M{"$in": bson.A{userID, bson.M{"$ifNull": bson.A{"$likes", bson.A{}}}}},
				"like",
				bson.M{
					"$cond": bson.A{
						bson.M{"$in": bson.A{userID, bson.M{"$ifNull": bson.A{"$dislikes", bson.A{}}}}},
						"dislike",
						"",
					},
				},
			},
		},
		// Include most fields
		"_id":         1,
		"value":       1,
		"itemId":      1,
		"itemType":    1,
		"title":       1,
		"description": 1,
		"userId":      1,
		"user":        1,
		"createdAt":   1,
		"updatedAt":   1,
		// TODO Add this as a function parameter to not query if not needed
		"commentIds": 1,
		"album":      1,
		// Exclude likes and dislikes
		"likes":    bson.M{"$cond": bson.A{true, bson.A{}, "$likes"}},    // Empty array
		"dislikes": bson.M{"$cond": bson.A{true, bson.A{}, "$dislikes"}}, // Empty array
	}
}

func GetCommentProjection(userID primitive.ObjectID) *bson.M {
	return &bson.M{
		// Calculate fields
		"likesCount":    bson.M{"$size": bson.M{"$ifNull": bson.A{"$likes", bson.A{}}}},
		"dislikesCount": bson.M{"$size": bson.M{"$ifNull": bson.A{"$dislikes", bson.A{}}}},
		"userReaction": bson.M{
			"$cond": bson.A{
				bson.M{"$in": bson.A{userID, bson.M{"$ifNull": bson.A{"$likes", bson.A{}}}}},
				"like",
				bson.M{
					"$cond": bson.A{
						bson.M{"$in": bson.A{userID, bson.M{"$ifNull": bson.A{"$dislikes", bson.A{}}}}},
						"dislike",
						"",
					},
				},
			},
		},
		// Include most fields
		"_id":       1,
		"reviewId":  1,
		"userId":    1,
		"text":      1,
		"createdAt": 1,
		"updatedAt": 1,
		// Exclude likes and dislikes
		"likes":    bson.M{"$cond": bson.A{true, bson.A{}, "$likes"}},    // Empty array
		"dislikes": bson.M{"$cond": bson.A{true, bson.A{}, "$dislikes"}}, // Empty array
	}
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
		projection = GetReviewProjection(userID)
	}
	if collection == "comments" {
		projection = GetCommentProjection(userID)
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
