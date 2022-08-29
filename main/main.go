package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/A-Victory/web/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.GhostWeb(r)
	fs := http.FileServer(http.Dir("static"))
	strp := http.StripPrefix(("/static/"), fs)
	r.PathPrefix("/static/").Handler(strp)
	http.Handle("/", r)
	fmt.Println("server is starting")
	log.Fatal(http.ListenAndServe(":80", r))

}
