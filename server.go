package gost

import (
	"gost/module"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
    routerHandler, err := s.Resolve()

    if err != nil {
        return err
    }

    
    return http.ListenAndServe(
        s.address, 
        routerHandler,
    )
}

func (s *BaseServer) Resolve() (http.Handler,error) {
    router := mux.NewRouter()

    for _, m := range s.modules {
        if err := m.ResolveModules(router); err != nil {
            return nil, err
        }
    }
	
    return router, nil
}
