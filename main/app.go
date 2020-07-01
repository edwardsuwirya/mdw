package main

import (
	"fmt"
	"github.com/edwardsuwirya/mdw/handler"
	"github.com/edwardsuwirya/mdw/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()
	arg := os.Args
	//Global Interceptor (Middleware)
	r.Use(middleware.ActivityLogMiddleware)

	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("", handler.NewAuthenticationHandler().Handler).Methods(http.MethodPost)
	auth.HandleFunc("/tokenValidation", handler.NewTokenValidationHandler().Handler).Methods(http.MethodGet)

	log.Printf("Server is listening on %s port %s", arg[1], arg[2])
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", arg[1], arg[2]), r); err != nil {
		log.Panic(err)
	}
}
