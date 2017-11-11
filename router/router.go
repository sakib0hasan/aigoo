package router

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
	"log"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func Router() {
	router := httprouter.New()
	port := "8080"
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	fmt.Println("Running on "+port+" port")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
