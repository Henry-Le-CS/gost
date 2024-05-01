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
		"/",
		getUserInfo,
	))

	router.Add(gost.DeclareRouteHandler(
		"GET",
		"/detail",
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