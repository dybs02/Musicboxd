package hlp

import (
	"fmt"
	"time"

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

var jwtSecret = Envs["JWT_SECRET"]

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

func ValidateJWT(tokenString string) (CustomClaims, error) {
	var claims CustomClaims

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return CustomClaims{}, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtSecret), nil
	})

	if err != nil {
		return claims, err
	}

	if !token.Valid {
		return claims, err
	}

	return claims, nil
}
