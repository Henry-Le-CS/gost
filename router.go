package gost

import (
	"github.com/gorilla/mux"
)

// ---------------------- Routers ----------------------
/*
	Declare a new router, in which there are many route handlers
*/
func DeclareRouter() *Router {
	return &Router{}
}

type Router struct {
	Handlers []*RouteHandler
}

// Add a route handler to the router
func (r * Router) Add(handler *RouteHandler) {
	// Add handler to handlers
	r.Handlers = append(r.Handlers, handler)
}

// Resolve all route handlers in the router
func (r * Router) Resolve(router * mux.Router) error {
	for _, h := range r.Handlers {
		if err := h.Resolve(router); err != nil {
			return err
		}
	}

	return nil
}