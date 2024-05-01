package exp

import (
	"net/http"

	"github.com/Henry-Le-CS/gost"
)

func TestControllers() *gost.Controller {
	router := gost.DeclareRouter()

	// Declare routes
	router.Add(gost.DeclareRouteHandler(
		"GET",
		"user",
		getUserInfo,
	))

	router.Add(gost.DeclareRouteHandler(
		"GET",
		"data",
		getUserData,
	))

	// Declare controllers
	controller := gost.DeclareController(gost.ControllerArgs{
		Prefix: "/user",
		Router: router,
	})

	return controller
}

// Define handlers
func getUserInfo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User info"))
}

// Define handlers
func getUserData(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User data"))
}