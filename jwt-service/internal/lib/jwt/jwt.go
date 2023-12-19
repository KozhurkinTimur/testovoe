package jwt

import (
	"jwt-service/internal/domain/entity"
	"log/slog"

	"github.com/golang-jwt/jwt/v5"
)

func NewToken(user *entity.User, log *slog.Logger) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	tokenString, err := token.SignedString([]byte(user.Secret))
	if err != nil {
		log.Error("error signed JWT token")
		return "", err
	}
	log.Info("generated new JWT token")
	return tokenString, nil
}
