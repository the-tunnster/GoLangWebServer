package main

import (
	"log"
	"net/http"
)

const portNumber = ":8080"

func main(){

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Println("Starting application on port", portNumber)

	http.ListenAndServe(portNumber, nil)

}