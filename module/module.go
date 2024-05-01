package module

import (
	"gost/controller"
	"log"
)

type Module struct {
	Name        string
	Controllers []controller.ControllerType
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

func (m *Module) GetControllers() []controller.ControllerType {
	return m.Controllers
}

func (m *Module) ResolveModules() error {
	// for _, c := range m.Controllers {
	// 	if err := c.Resolve(); err != nil {
	// 		return err
	// 	}
	// }
	return nil
}