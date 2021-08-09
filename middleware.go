package dependancy_middlewares

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/cmatikdev/middleware-go-dependancy/functions"
	"github.com/cmatikdev/middleware-go-dependancy/responses"
	"github.com/gorilla/mux"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := functions.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}

func GetMiddlewareRol(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid, err := functions.ExtractTokenID(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized no token"))
			return
		}
		currentRoute := mux.CurrentRoute(r)
		if currentRoute != nil {
			fmt.Printf("uuid: %v\n", uuid)
			fmt.Printf("currentRoute.GetName(): %v\n", currentRoute.GetName())
			val, _ := functions.ValidateRole(uuid, currentRoute.GetName())
			if !val {
				responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized permisssions"))
				return
			}
		}
		next(w, r)
	}
}
