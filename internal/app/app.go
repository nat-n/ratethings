package app

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)


func Run() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/echo", Echo)
	log.Fatal(http.ListenAndServe(":8080", router))
}


func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func Echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Echo, %q", html.EscapeString(r.URL.Path))
}
