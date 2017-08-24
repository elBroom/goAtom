package server

import (
	"net/http"
	"github.com/gorilla/mux"

	"app/router"
	"app/workers"
)

func init() {
	workers.Wp.Run()
}

func RunHTTPServer(addr string) error {
	_router := mux.NewRouter()

	router.Routing(_router)
	return http.ListenAndServe(addr, _router)
}
