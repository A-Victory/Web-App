package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/A-Victory/web/pkg/routes"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()

	routes.GhostWeb(r)
	/*
		fs := http.FileServer(http.Dir("static"))
		strp := http.StripPrefix(("/static/"), fs)
	*/
	cert, _ := tls.LoadX509KeyPair("web.crt", "web.key")
	r.ServeFiles("/static/*filepath", http.Dir("static"))
	fmt.Println("server is starting")
	srv := &http.Server{
		Addr:    ":9000",
		Handler: r,

		TLSConfig: &tls.Config{Certificates: []tls.Certificate{cert}},
	}

	log.Fatal(srv.ListenAndServeTLS("", ""))

}
