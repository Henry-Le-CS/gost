package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

// ---------------------- RouteHandler ----------------------
type RouteHandler struct {
	method string
	pattern string
	handler func(w http.ResponseWriter, r *http.Request)
}

type RouteHandlerError struct {
	message string
}

func (e *RouteHandlerError) Error() string {
	return e.message
}

func DeclareRouteHandler(method string, pattern string, handler func(w http.ResponseWriter, r *http.Request)) *RouteHandler {
	if method == "" {
		method = "GET"
	}

	if pattern == "" {
		pattern = "/"
	}

	return &RouteHandler{method: method, pattern: pattern, handler: handler}
}

func (rh * RouteHandler) Resolve(router * mux.Router) error {
	r := router.HandleFunc(rh.pattern, rh.handler).Methods(rh.method)

	if r == nil {
		return &RouteHandlerError{message: "Failed to resolve route handler"}
	}

	return nil
}

// ---------------------- Routers ----------------------
/*
	Declare a new router, in which there are many route handlers
*/

func DeclareRouter() IRouter {
	return &Router{}
}

type Router struct {
	Handlers []*RouteHandler
}

func (r * Router) Add(handler *RouteHandler) {
	// Add handler to handlers
	r.Handlers = append(r.Handlers, handler)
}

func (r * Router) Resolve(router * mux.Router) error {
	for _, h := range r.Handlers {
		if err := h.Resolve(router); err != nil {
			return err
		}
	}

	return nil
}