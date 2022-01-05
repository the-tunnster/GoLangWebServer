package main

import (
	"WebServer/pkg/config"
	"WebServer/pkg/handlers"
	"WebServer/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main(){

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if(err!=nil){
		log.Fatal("Cannot create template cache : ", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	log.Println("Starting application on port", portNumber)

	err = srv.ListenAndServe()
	if(err!=nil){
		log.Fatal(err)
	}
	
}