package module

import "gost/controller"

type RegisterModulePayload struct {
	name string
	controllers []controller.ControllerType
}