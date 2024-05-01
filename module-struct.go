package gost

type RegisterModuleDto struct {
	Name string
	Controllers []IController
}