package resolver

import (
	"context"
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
)

//go:generate go run github.com/99designs/gqlgen generate
// This file will not be regenerated automatically.
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

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
		"value":       1,
		"itemId":      1,
		"itemType":    1,
		"title":       1,
		"description": 1,
		"userId":      1,
		"user":        1,
		"createdAt":   1,
		"updatedAt":   1,
		"comments":    1,
		"album":       1,
		// Exclude likes and dislikes
		"likes":    bson.M{"$cond": bson.A{true, bson.A{}, "$likes"}},    // Empty array
		"dislikes": bson.M{"$cond": bson.A{true, bson.A{}, "$dislikes"}}, // Empty array
	}
}
