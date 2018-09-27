package MVC

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

type View struct {

}

func (v View) RenderTemplate(w http.ResponseWriter, tmpl string, p *Model) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}