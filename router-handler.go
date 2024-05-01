package gost

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
	return "[Route handler]: " + e.message
}

// Declare RouteHandler declare a new route handler
func DeclareRouteHandler(method string, pattern string, handler func(w http.ResponseWriter, r *http.Request)) *RouteHandler {
	if method == "" {
		method = "GET"
	}

	if pattern == "" {
		pattern = "/"
	}

	return &RouteHandler{method: method, pattern: pattern, handler: handler}
}

// Resolve the route handler
func (rh * RouteHandler) Resolve(router * mux.Router) error {
	r := router.HandleFunc(rh.pattern, rh.handler).Methods(rh.method)

	if r == nil {
		msg := "Failed to resolve route handler for pattern: " + rh.pattern + " and method: " + rh.method
		return &RouteHandlerError{message: msg}
	}

	return nil
}
