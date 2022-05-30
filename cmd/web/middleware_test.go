package main

import ("testing"
		"net/http"
		"fmt")

func TestNoSurf(t *testing.T){

	var myH myHandler
	h := NoSurf(&myH)

	switch v := h.(type){
	case http.Handler:
		//Doing nothing
	default:
		t.Error(fmt.Sprintf("Type is not http.Handler, but is %t.", v))
	}

}

func TestSessionLoad(t *testing.T){

	var myH myHandler
	h := SessionLoad(&myH)

	switch v := h.(type){
	case http.Handler:
		//Doing nothing
	default:
		t.Error(fmt.Sprintf("Type is not http.Handler, but is %t.", v))
	}

}