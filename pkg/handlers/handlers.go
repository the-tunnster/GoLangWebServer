package handlers

import (
	"WebServer/pkg/config"
	"WebServer/pkg/render"
	"WebServer/pkg/models"
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
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request){

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}