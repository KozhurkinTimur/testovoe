package jwt

import (
	"testovoe/internal/models"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

func NewToken(authorization *models.Authorization) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = authorization.Id
	tokenString, err := token.SignedString([]byte(authorization.Secret))
	if err != nil {
		log.Fatal("error signed JWT token")
		return "", err
	}
	//log.Info("generated new JWT token")
	return tokenString, nil
}
