package handlers

import (
	"WebServer/pkg/config"
	"WebServer/pkg/models"
	"WebServer/pkg/render"
	"net/http"
)

//Is the the repository type
type Repository struct{
	App *config.AppConfig
}

var Repo *Repository

//Creates a new repository
func NewRepo(a *config.AppConfig) *Repository{
	return &Repository{
		App : a,
	}
}

//Sets the repository for handlers
func NewHandlers(r *Repository){
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request){

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})

}

func (m *Repository) About(w http.ResponseWriter, r *http.Request){

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}