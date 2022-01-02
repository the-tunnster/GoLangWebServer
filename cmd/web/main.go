package main

import (
	"WebServer/pkg/handlers"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main(){

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Println("Starting application on port", portNumber)

	http.ListenAndServe(portNumber, nil)

}