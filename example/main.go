package main

import (
	"fmt"

	"github.com/Henry-Le-CS/gost"
	"github.com/Henry-Le-CS/gost/example/exp"
)

func main() {
	modules := []gost.IModule{exp.TestModule()}

	s := gost.NewServer("localhost:8080", modules, &gost.ServerOptions{
		CorsOptions: gost.CorsOptions{
			AllowedOrigins: []string{"*"},
			AllowCredentials: true,
		},
	})

	if err := s.Start(); err != nil {
		fmt.Println(err)
	}
}