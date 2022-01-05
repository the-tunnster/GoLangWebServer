package main

import (
	"WebServer/pkg/config"
	"WebServer/pkg/handlers"
	"WebServer/pkg/render"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main(){

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

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