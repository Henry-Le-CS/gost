package gost

import "github.com/gorilla/mux"

type ControllerArgs struct {
	Prefix string
	Router IRouter
}

// Declare set a controller with a sets of routers
func DeclareController(args ControllerArgs) *Controller {
    return &Controller{router: args.Router, prefix: args.Prefix}
}

type Controller struct {
	router IRouter
	prefix string
}

// Resolve resolve all routers in the controller
func (c *Controller) Resolve(r *mux.Router) error {
	if c.prefix != "" {
		r = r.PathPrefix(c.prefix).Subrouter()
	}

	if err := c.router.Resolve(r); err != nil {
		return err
	}

	return nil
}