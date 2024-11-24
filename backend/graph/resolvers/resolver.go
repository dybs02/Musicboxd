package resolver

import (
	"context"
	"fmt"
	"musicboxd/api/middleware"
	"musicboxd/endpoints/auth"
	"musicboxd/hlp"

	"github.com/gin-gonic/gin"
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
