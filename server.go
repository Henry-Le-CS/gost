package gost

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type CorsOptions struct {
    AllowedOrigins []string
    AllowCredentials bool
}

type ServerOptions struct {
    CorsOptions CorsOptions
}

type BaseServer struct {
    address  string
    modules  []IModule
    options ServerOptions
}

func NewServer(address string, modules []IModule, options *ServerOptions) *BaseServer {
    return &BaseServer{
        address: address,
        modules: modules,
        options: *options,
    }
}

func (s *BaseServer) Start() error {
    log.Printf("Server is running at %s", s.address)
    routerHandler, err := s.Resolve()

    if err != nil {
        return err
    }

   if s.options.CorsOptions.AllowedOrigins != nil {
        c := cors.New(cors.Options{
            AllowedOrigins: s.options.CorsOptions.AllowedOrigins,
            AllowCredentials: s.options.CorsOptions.AllowCredentials,
        })

        routerHandler = c.Handler(routerHandler)
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
