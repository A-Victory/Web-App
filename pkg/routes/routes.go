package routes

import (
	controller "github.com/A-Victory/web/pkg/controllers"
	"github.com/gorilla/mux"
)

var GhostWeb = func(router *mux.Router) {
	router.HandleFunc("/home", controller.Home)
	router.HandleFunc("/contact", controller.Contact)
	router.HandleFunc("/about", controller.About)
	router.HandleFunc("/projects", controller.Project)
	router.HandleFunc("/tests", controller.Testimonials)
}
