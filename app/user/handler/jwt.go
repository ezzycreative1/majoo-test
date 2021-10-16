package handler

import (
	"fmt"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			//nil secret key
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
}

func ParsingTokentoUserID(token string) string {
	claims := jwt.MapClaims{}
	result, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("API_SECRET")), nil

	})

	if err != nil {
		fmt.Println(err)
	}
	log.Print(result)
	return claims["userID"].(string)

}

func GetUserID(ctx *gin.Context) string {
	const BearerSchema string = "Bearer "
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len(BearerSchema):]
	userID := ParsingTokentoUserID(tokenString)
	return userID
}
