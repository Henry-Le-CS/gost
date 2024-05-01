package controller

import (
	"github.com/gorilla/mux"
)

type Controller struct {
	name string
	handlers []RouteHandler
}



func DeclareController(name string, handlers []RouteHandler) *Controller {
	return &Controller{name: name, handlers: handlers}
}

func (c *Controller) Resolve(router * mux.Router) error {
	for _, h := range c.handlers {
		if err := h.Resolve(router); err != nil {
			return err
		}
	}

	return nil
}