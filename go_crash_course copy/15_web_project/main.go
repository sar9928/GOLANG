package main

import (
	"fmt"
	"log"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("website").HTTPBox()))
	fmt.Println("Server Starting...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
