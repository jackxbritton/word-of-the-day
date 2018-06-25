package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handleMessenger(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("hi"))

}

func main() {

	port := flag.Int("port", 8080, "port to listen on")
	flag.Parse()

	router := mux.NewRouter()
	router.HandleFunc("/messenger", handleMessenger)

	http.ListenAndServe(fmt.Sprintf(":%d", *port), router)

}
