package module

import "gost/controller"

type RegisterModuleDto struct {
	Name string
	Controllers []controller.ControllerType
}