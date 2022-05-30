package main

import (
	"WebServer/internal/config"
	"fmt"
	"testing"

	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T){
	var app config.AppConfig

	mux := routes(&app)
	
	switch v:= mux.(type){
	case *chi.Mux:
		//do nothing
	default:
		t.Error(fmt.Sprintf("Type is not chi.Mux, but is %t.", v))
	}
}