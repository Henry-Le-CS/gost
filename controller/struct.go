package controller

import "net/http"

type ControllerType struct {
	method string
	route string
	handler http.Handler
}
