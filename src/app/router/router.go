package router

import (
	"github.com/gorilla/mux"

	"../handler"
)

func Routing(router *mux.Router) {
	router.HandleFunc("/value/", handler.CreateValueEndpoint).Methods("POST")
	router.HandleFunc("/value/{key}", handler.GetValueEndpoint).Methods("GET")
	router.HandleFunc("/value/{key}", handler.UpdateValueEndpoint).Methods("PUT")
	router.HandleFunc("/value/{key}", handler.DeleteValueEndpoint).Methods("DELETE")

	router.HandleFunc("/user", handler.CreateUserEndpoint).Methods("POST")
	router.HandleFunc("/login", handler.AuthUserQuery).Methods("POST")
	router.HandleFunc("/login", handler.LogoutUserQuery).Methods("POST")

	router.HandleFunc("/history", handler.GetHistoryEndpoint).Methods("GET")
}
