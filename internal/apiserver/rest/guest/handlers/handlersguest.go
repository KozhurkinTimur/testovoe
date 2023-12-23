package handlers

import (
	"log/slog"
	"net/http"
	"testovoe/internal/lib/jwt"
	"testovoe/internal/models"
)

func handlerGuest(w http.ResponseWriter, r *http.Request, log *slog.Logger) (string, error){
 	token, err := jwt.NewToken(&models.Authorization{}, log)
	return token, err
}