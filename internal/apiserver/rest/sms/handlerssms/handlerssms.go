package handlerssms

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testovoe/internal/storage/postgresql"

	"github.com/golang-jwt/jwt/v5"
)

type SMSHandler struct {
	Id int
}

func New() *SMSHandler {
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

	if err != nil {
		// Обработка ошибок
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Токен декодирован и валидный
		requestBody, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		// Декодировать JSON тело запроса в map[string]interface{}
		var bodyMap map[string]interface{}
		err = json.Unmarshal(requestBody, &bodyMap)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	
		// Получить значения из map
		login, ok1 := bodyMap["login"].(string)
		password, ok2 := bodyMap["password"].(string)
		
		if !ok1 || !ok2 {
			http.Error(w, "Invalid field types", http.StatusBadRequest)
			return
		}

		uid := claims["uid"]
		db, err = postgresql.Connect()
		sql := &postgresql.PostgresUserRepository{}
	}
	
	func (r *PostgresUserRepository) Connect() (*sql.DB, error) {
		//db, err := sql.Open("
		sql.Save(login, password, uid.(string))

		w.Write([]byte("Uid: " + fmt.Sprint(uid)))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}

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
