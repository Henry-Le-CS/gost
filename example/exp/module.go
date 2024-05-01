package exp

import "github.com/Henry-Le-CS/gost"

func TestModule() *gost.Module {
	controllers := []gost.IController{TestControllers()}

	module := gost.DeclareModule(gost.RegisterModuleDto{
		Name: "Test controller",
		Controllers: controllers,
	})

	return module
}