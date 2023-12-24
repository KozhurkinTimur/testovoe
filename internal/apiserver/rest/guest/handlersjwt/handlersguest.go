package handlersjwt

import (
	"fmt"
	"log/slog"
	"net/http"
	"testovoe/internal/lib/jwt"
	"testovoe/internal/models"
)

type JWTHandler struct {
	Id int
}

func New() *JWTHandler {
	return &JWTHandler{
		Id: 3,
	}
}

func (*JWTHandler) HandlerGuest(w http.ResponseWriter, r *http.Request) {
 	token, err := jwt.NewToken(&models.Authorization{}, &slog.Logger{})
	fmt.Fprintf(w, "Token: %s, %s", token, err)
}