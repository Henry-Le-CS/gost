package server

import (
	"gost/module"
	"log"
	"net/http"
)

type Server struct {
    address string
	modules []module.Module
}

func NewServer(address string, modules []module.Module) *Server {
    return &Server{
        address: address,
        modules: modules,
    }
}

func (s *Server) Start() error {
	log.Printf("Server is running at %s", s.address)
    return http.ListenAndServe(s.address, nil)
}

// TODO: resolve modules and routers
func (s* Server) ResolveModules() {
	for _, module := range s.modules {
		for _, controller := range module.GetControllers() {
			// http.Handle(controller, controller.handler)
		}
	}
}
