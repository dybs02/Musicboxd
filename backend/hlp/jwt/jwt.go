package jwt

import (
	"context"
	"fmt"
	"musicboxd/api/middleware"
	"musicboxd/hlp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomClaims struct {
	UserID primitive.ObjectID `json:"id"`
	Email  string             `json:"email"`
	Role   string             `json:"role"`
	jwt.StandardClaims
}

type UserInfo struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

var jwtSecret = hlp.Envs["JWT_SECRET"]

func GenerateJWT(userInfo UserInfo) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	convertedID, err := primitive.ObjectIDFromHex(userInfo.ID)
	if err != nil {
		return "", err
	}

	claims := CustomClaims{
		UserID: convertedID,
		Email:  userInfo.Email,
		Role:   userInfo.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
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

func ValidateJWT(tokenString string) (*CustomClaims, error) {
	var claims CustomClaims

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return CustomClaims{}, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtSecret), nil
	})

	if err != nil {
		return &claims, err
	}

	if !token.Valid {
		return &claims, err
	}

	return &claims, nil
}

func ValidateJWTFromCtx(ctx context.Context) (*CustomClaims, error) {
	ginCtx, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	tokenString := ginCtx.GetHeader("Authorization")
	if tokenString == "" {
		return nil, fmt.Errorf("authorization header is missing")
	}

	return ValidateJWT(tokenString)
}
