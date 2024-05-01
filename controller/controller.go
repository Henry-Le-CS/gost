package controller

import "github.com/gorilla/mux"


func DeclareController(routers IRouter) *Controller {
	rt := make([]IRouter, 0)
	rt = append(rt, routers)

    return &Controller{routers: rt}
}

type Controller struct {
	routers []IRouter
}

func (c *Controller) AddRouter(router IRouter) {
	c.routers = append(c.routers, router)
}

func (c *Controller) Resolve(r *mux.Router) error {
	for _, router := range c.routers {
		if err := router.Resolve(r); err != nil {
			return err
		}
	}
	return nil
}