package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mastama/go-jwt-mux/controller/authcontroller"
	productcontroller "github.com/mastama/go-jwt-mux/controller/authcontroller/productController"
	"github.com/mastama/go-jwt-mux/model"

	"github.com/mastama/go-jwt-mux/middleware"
)

func main() {

	model.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("logout", authcontroller.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/products", productcontroller.Index).Methods("GET")
	api.Use(middleware.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":8081", r))

}
