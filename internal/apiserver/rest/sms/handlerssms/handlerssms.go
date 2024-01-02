package handlerssms

import (
	"encoding/json"
	"net/http"
	"testovoe/internal/models"
	"testovoe/internal/storage"

	"github.com/golang-jwt/jwt/v5"
)

type RegReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type SMSHandler struct {
	repo storage.UserRepository
}

func New(repo storage.UserRepository) *SMSHandler {
	return &SMSHandler{
		repo: repo,
	}
}

func (h *SMSHandler) HandlerRegistration(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("sasa"), nil
	})

	if err != nil || !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims);
	uid := claims["uid"]

	if !ok || uid == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var req RegReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.repo.Save(&models.User{
		Id: uid.(string),
		Login:    req.Login,
		Password: req.Password,
	})

}


func (h *SMSHandler) HandlerSignIn(w http.ResponseWriter, r *http.Request) {}

func (h *SMSHandler) HandlerPay(w http.ResponseWriter, r *http.Request) {}

func (h *SMSHandler) HandlerConfirmations(w http.ResponseWriter, r *http.Request) {}
