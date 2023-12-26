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
	Id   int
}

func New(repo storage.UserRepository) *SMSHandler {
	return &SMSHandler{
		Id: 2,
	}
}

func (h *SMSHandler) HandlerRegistration(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	// Убедитесь, что в заголовке "Authorization" ваш токен выглядит примерно так: "Bearer <ваш токен>"
	// Здесь нужно разобрать токен из заголовка запроса и декодировать его
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ваш код здесь для извлечения секретного ключа для проверки подписи токена
		return []byte("sasa"), nil
	})

	if err != nil || !token.Valid {
		// Обработка ошибок
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var req RegReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.repo.Save(&models.User{
		Login:    req.Login,
		Password: req.Password,
	})

}

// TODO: verify uuid in db, registartion id in another provider and db
// TODO: verify correctly login and password
// TODO: send sms with confirmation code
// TODO: create user in db

func (h *SMSHandler) HandlerSignIn(w http.ResponseWriter, r *http.Request) {
	// TODO: verify uuid in db, login and password
	// TODO: verify confirmation code
	// TODO: login user
}

func (h *SMSHandler) HandlerPay(w http.ResponseWriter, r *http.Request) {
	// TODO: get number, key uuid, some data
	// TODO: send sms with confirmation code
}

func (h *SMSHandler) HandlerConfirmations(w http.ResponseWriter, r *http.Request) {
	// TODO: get uuid and code
	// TODO: verify code and uuid
}
