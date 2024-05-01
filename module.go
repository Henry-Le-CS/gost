package gost

import (
	"log"

	"github.com/gorilla/mux"
)

type Module struct {
	Name        string
	Controllers []IController
}

func DeclareModule(payload RegisterModuleDto) *Module {
	m := &Module{}
	return m.Register(payload)
}

func (m *Module) Register(payload RegisterModuleDto) *Module {
	m.Controllers = payload.Controllers
	m.Name = payload.Name

	log.Printf("Module %s is registered", m.Name)
	return m
}

/*
	Resolve all controllers in the module
*/
func (m *Module) ResolveModules(router *mux.Router) error {
	for _, c := range m.Controllers {
		if err := c.Resolve(router); err != nil {
			return err
		}
	}

	return nil
}