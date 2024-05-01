package server

import (
	"gost/module"
	"log"
	"net/http"
)

type BaseServer struct {
    address  string
    modules  []module.IModule
}

func NewServer(address string, modules []module.IModule) *BaseServer {
    return &BaseServer{
        address: address,
        modules: modules,
    }
}

func (s *BaseServer) Start() error {
    log.Printf("Server is running at %s", s.address)
    return http.ListenAndServe(s.address, nil)
}

func (s *BaseServer) ResolveModules() error {
    for _, m := range s.modules {
        if err := m.ResolveModules(); err != nil {
            return err
        }
    }
	
    return nil
}
