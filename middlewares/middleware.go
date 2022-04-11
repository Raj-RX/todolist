package middlewares

import (
	"context"
	"github.com/tejashwikalptaru/tutorial/database/helper"
	"github.com/tejashwikalptaru/tutorial/models"
	"net/http"
)

type ContextKeys string

const (
	userContext ContextKeys = "__userContext"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Code for the middlewares...

		apiKey := r.Header.Get("Authorization")

		user, err := helper.GetUserBySession(apiKey)

		if err != nil || user == nil {
			//logrus.WithError(err).Errorf("failed to get user with token: %s", apiKey)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userContext, user)
		next.ServeHTTP(w, r.WithContext(ctx))

		//next.ServeHTTP(w, r)
	})
}

func UserContext(r *http.Request) *models.User {
	if user, ok := r.Context().Value(userContext).(*models.User); ok && user != nil {
		return user
	}
	return nil
}

//var token models.Session
//err := json.NewDecoder(request.Body).Decode(&token)
//if err != nil {
//	log.Fatal(err)
//}
//SQL := `Select log_out from sessions where id = $1`
//var tok models.Session
//err = database.Tutorial.Get(&tok, SQL, token.SessioId)
//if err != nil {
//	panic(err)
//}
//fmt.Println(tok.Logut)
//token := request.Header["Authorization"]
//
//SQL := `SELECT id, uid, login_at, log_out from sessions where id = $1`
//var userLoginStatus models.Session
//err := database.Tutorial.Get(&userLoginStatus, SQL, token)
//if err != nil {
//	writer.WriteHeader(400)
//	fmt.Println(token)
//	return
//}
//err = json.NewEncoder(writer).Encode(userLoginStatus)
//fmt.Println(userLoginStatus.SessioId , userLoginStatus.UserId , userLoginStatus.Login , userLoginStatus.Logut)
