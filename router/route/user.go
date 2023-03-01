package route

import (
	"github.com/controller"
	"net/http"
)

const userResource = "/users"
const userResourceByID = userResource + "/{id}"

var userRoutes = []Route{
	{
		URI:             userResource,
		Method:          http.MethodGet,
		Handler:         controller.FindAllUsers,
		IsAuthenticated: false,
	},
	{
		URI:             userResourceByID,
		Method:          http.MethodGet,
		Handler:         controller.FindUserByID,
		IsAuthenticated: false,
	},
	{
		URI:             userResource,
		Method:          http.MethodPost,
		Handler:         controller.CreateUser,
		IsAuthenticated: false,
	},
	{
		URI:             userResourceByID,
		Method:          http.MethodDelete,
		Handler:         controller.DeleteUserByID,
		IsAuthenticated: false,
	},
	{
		URI:             userResourceByID,
		Method:          http.MethodPut,
		Handler:         controller.UpdateUserByID,
		IsAuthenticated: false,
	},
}
