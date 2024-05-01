package controller

import "github.com/gorilla/mux"

type IController interface {
	DeclareController(name string, handlers []RouteHandler) *Controller
	Resolve(router * mux.Router) error
}