package handlersjwt

import (
	"encoding/json"
	"fmt"

	"net/http"
	"testovoe/internal/config"
	"testovoe/internal/lib/jwt"
	"testovoe/internal/models"
)

type JWTHandler struct {}

func New() *JWTHandler {
	return &JWTHandler{}
}

func (*JWTHandler) HandlerSignIn(w http.ResponseWriter, r *http.Request) {
	token, err := jwt.NewToken(&models.Authorization{})
	if err != nil {
		fmt.Fprintf(w, "Error generate Token for Guest:  %s", err)
	}
	// res := map[string]interface{}{
	// 	"access_token": token,
	// }
	data, _ := json.Marshal(token)
	fmt.Fprintf(w, string(data))
}

func (*JWTHandler) HandlerGuest(w http.ResponseWriter, r *http.Request) {

	token, err := jwt.NewToken(models.New(config.MustLoad()))
	if err != nil {
		fmt.Fprintf(w, "Error generate Token for Guest:  %s", err)
	}
	// res := map[string]interface{}{
	// 	"access_token": token,
	// }
	data, _ := json.Marshal(token)
	fmt.Fprintf(w, string(data))

}
