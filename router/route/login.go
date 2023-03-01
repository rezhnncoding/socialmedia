package route

import (
	"github.com/controller"
	"net/http"
)

var loginRoute = Route{
	URI:             "/login",
	Method:          http.MethodPost,
	Handler:         controller.Login,
	IsAuthenticated: false,
}
