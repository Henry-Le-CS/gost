package module

import (
	"gost/controller"
	"log"
)

type Module struct {
	name string
	controllers []controller.ControllerType
}

func (m *Module) Register(payload RegisterModulePayload) *Module {
	m.controllers = payload.controllers
	m.name = payload.name

	log.Printf("Module %s is registered", m.name)
	return m
}

func (m *Module) GetControllers() []controller.ControllerType {
	return m.controllers
}

// TODO: resolve dependencies
// TODO: resolve controllers