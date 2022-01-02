package render

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
)

var functions = template.FuncMap{

}

func RenderTemplate(w http.ResponseWriter, tmpl string){

	
	_, err := RenderTemplateTest(w)
	if(err!=nil){
		fmt.Println("Error getting template cahce : ", err)
	}

	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err = parsedTemplate.Execute(w, nil)
	if(err!=nil){
		fmt.Println("Error parsing template : ", err)
	}
	
}

func RenderTemplateTest(w http.ResponseWriter) (map[string]*template.Template, error){

	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if(err!=nil){
		return myCache, err
	}

	for _, page := range pages{

		name := filepath.Base(page)
		fmt.Println("Page is currently : ", page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if(err!=nil){
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if(err!=nil){
			return myCache, err
		}

		if(len(matches)>0){
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if(err!=nil){
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil

}