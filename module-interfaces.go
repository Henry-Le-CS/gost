package gost

import "github.com/gorilla/mux"

type IModule interface {
	Register(payload RegisterModuleDto) *Module
	ResolveModules(router *mux.Router) error
}