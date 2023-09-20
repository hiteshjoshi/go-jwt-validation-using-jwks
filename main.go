package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/hiteshjoshi/jwt-validation/jwt"
	"github.com/hiteshjoshi/jwt-validation/middleware"
)

func main() {

	httpClient := jwt.JWKHttpClient{}
	verifier := jwt.JwtTokenVerifier{
		JWKSUri:    os.Getenv("JWKS_URI"),
		HTTPClient: &httpClient,
	}

	r := mux.NewRouter()

	//200 on success
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	amw := middleware.Auth{
		JwtTokenVerifier: verifier,
	}
	//middleware to validate jwt
	r.Use(amw.Middleware)

	err := http.ListenAndServe(":3000", r)
	log.Fatal(err)
}
