package jwt

import (
	"jwt-service/internal/models"
	"log/slog"

	"github.com/golang-jwt/jwt/v5"
)

func NewToken(authorization *models.Authorization, log *slog.Logger) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = authorization.ID
	tokenString, err := token.SignedString([]byte(authorization.Secret))
	if err != nil {
		log.Error("error signed JWT token")
		return "", err
	}
	log.Info("generated new JWT token")
	return tokenString, nil
}
