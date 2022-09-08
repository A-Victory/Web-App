package routes

import (
	controller "github.com/A-Victory/web/pkg/controllers"
	"github.com/julienschmidt/httprouter"
)

var GhostWeb = func(router *httprouter.Router) {
	router.GET("/home", controller.Home)
	router.GET("/contact", controller.Contact)
	router.GET("/about", controller.About)
	router.GET("/projects", controller.Project)
	router.GET("/tests", controller.Testimonials)
}
