package route

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/middleware"
	"net/http"
)

type Route struct {
	URI             string
	Method          string
	Handler         func(http.ResponseWriter, *http.Request)
	IsAuthenticated bool
}

func AddRoutes(r *mux.Router) *mux.Router {
	routes := append(userRoutes, loginRoute)

	fmt.Println("Routes creation started.")

	for _, route := range routes {
		addRoute(r, route)
	}

	fmt.Println("Routes creation completed.")

	return r
}

func addRoute(r *mux.Router, route Route) {
	r.HandleFunc(route.URI, getHandleFunc(route)).Methods(route.Method)
	fmt.Printf("Route [%s]	%s added.\n", route.Method, route.URI)
}

func getHandleFunc(route Route) http.HandlerFunc {
	if route.IsAuthenticated {
		return middleware.Authenticate(route.Handler)
	}

	return route.Handler
}
