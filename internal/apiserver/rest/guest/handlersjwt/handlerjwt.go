package handlersjwt

import (
	"encoding/json"
	"fmt"

	"net/http"
	"testovoe/internal/config"
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

func (*JWTHandler) HandlerSignIn(w http.ResponseWriter, r *http.Request) {
	token, err := jwt.NewToken(&models.Authorization{})
	if err != nil {
		fmt.Fprintf(w, "Error generate Token for Guest:  %s", err)
	}
	res := map[string]interface{}{
		"access_token": token,
	}
	data, _ := json.Marshal(res)
	fmt.Fprintf(w, string(data))
}

func (*JWTHandler) HandlerGuest(w http.ResponseWriter, r *http.Request) {

	token, err := jwt.NewToken(models.New(config.MustLoad()))
	if err != nil {
		fmt.Fprintf(w, "Error generate Token for Guest:  %s", err)
	}
	res := map[string]interface{}{
		"access_token": token,
	}
	data, _ := json.Marshal(res)
	fmt.Fprintf(w, string(data))

}

// func LogHandler(fn http.HandlerFunc) http.HandlerFunc {
//     return func(w http.ResponseWriter, r *http.Request) {
//         x, err := httputil.DumpRequest(r, true)
//         if err != nil {
//             http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
//             return
//         }
//         log.Println(fmt.Sprintf("%q", x))
//         rec := httptest.NewRecorder()
//         fn(rec, r)
//         log.Println(fmt.Sprintf("%q", rec.Body))
//     }
// }

// func MessageHandler(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintln(w, "A message was received")
// }
