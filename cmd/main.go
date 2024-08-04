package main

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Auth struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}

var validate *validator.Validate

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	auth := &Auth{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	authErr := validate.Struct(auth)
	if authErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	auth := &Auth{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	authErr := validate.Struct(auth)
	if authErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	validate = validator.New(validator.WithRequiredStructEnabled())

	http.Handle("/auth/login", corsMiddleware(http.HandlerFunc(loginHandler)))
	http.Handle("/auth/register", corsMiddleware(http.HandlerFunc(registerHandler)))

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
