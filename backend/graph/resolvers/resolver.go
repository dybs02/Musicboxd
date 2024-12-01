package resolver

import (
	"context"
	"fmt"
	"musicboxd/api/middleware"
	"musicboxd/database"
	"musicboxd/endpoints/auth"
	"musicboxd/graph/model"
	"musicboxd/hlp"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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

	jwt, err := ginCtx.Cookie(auth.JWT_KEY)
	if err != nil {
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
