package router

import (
	"github.com/gorilla/mux"
	"github.com/router/route"
)

func Build() *mux.Router {
	r := mux.NewRouter()
	return route.AddRoutes(r)
}
