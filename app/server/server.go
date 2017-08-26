package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/elBroom/goAtom/app/router"
	"github.com/elBroom/goAtom/app/workers"
)

func init() {
	workers.Wp.Run()
}

func RunHTTPServer(addr string) error {
	_router := mux.NewRouter()

	_router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	router.Routing(_router)
	return http.ListenAndServe(addr, _router)
}
