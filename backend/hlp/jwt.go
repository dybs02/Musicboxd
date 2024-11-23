package hlp

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	UserID string `json:"id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

type UserInfo struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

var jwtSecret = Envs["JWT_SECRET"]

func GenerateJWT(userInfo UserInfo) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := CustomClaims{
		UserID: userInfo.ID,
		Email:  userInfo.Email,
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
