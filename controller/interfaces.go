package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

type IController interface {
	AddRouter(router IRouter)
	Resolve(router * mux.Router) error
}

type IRouterHandler interface {
	DeclareRouteHandler(method string, pattern string, handler func(w http.ResponseWriter, r *http.Request)) *RouteHandler
}

type IRouter interface {
	Add(handler *RouteHandler)
	Resolve(router * mux.Router) error
}