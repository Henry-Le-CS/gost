package module

import "gost/controller"

type IModule interface {
	Register(payload RegisterModuleDto) *Module
	GetControllers() []controller.ControllerType
	ResolveModules() error
}