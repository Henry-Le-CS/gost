package gost

import "gost/controller"

type RegisterModuleDto struct {
	Name string
	Controllers []controller.IController
}