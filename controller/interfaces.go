package controller

import (
	"github.com/gorilla/mux"
)

type IController interface {
	Resolve(router * mux.Router) error
}

type IRouter interface {
	Add(handler *RouteHandler)	
	Resolve(router * mux.Router) error
}