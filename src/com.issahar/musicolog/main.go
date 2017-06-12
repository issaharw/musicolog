package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
	"log"
)


var router = httprouter.New()


func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello!")
}


func main() {
	router.GET("/hello", Hello)

	log.Fatal(http.ListenAndServe(":8080", router))
}


