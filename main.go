package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is the home page.")
}

func About(w http.ResponseWriter, r *http.Request){

	sum := AddValues(6, 9)
	fmt.Fprintf(w, "This is the about page.\n")
	fmt.Fprintf(w, "The sum of two values is %d.", sum)

}

func AddValues(x, y int) int {	return(x+y)		}

func main(){

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("Starting application on port", portNumber)

	http.ListenAndServe(portNumber, nil)

}