package controller

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
	//"path/filepath"
)

//const tp = "static"

//var pattern = filepath.Join(tp, "*.html")

var Temp *template.Template

func init() {
	Temp = template.Must(template.ParseGlob("assets/*.html"))
}

func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	//fmt.Fprint(w, "Welcome to my site")
	Temp.ExecuteTemplate(w, "index.html", nil)
}

func Contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	Temp.ExecuteTemplate(w, "contact.html", nil)
}

func About(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	Temp.ExecuteTemplate(w, "about.html", nil)
}

func Project(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	Temp.ExecuteTemplate(w, "projects.html", nil)
}

func Testimonials(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	Temp.ExecuteTemplate(w, "testimonials.html", nil)
}
