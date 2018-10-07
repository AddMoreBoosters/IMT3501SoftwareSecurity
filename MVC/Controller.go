package MVC 

import (
	"net/http"
	"regexp"
	"log"
)

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

type Controller struct {
	model Model
	view View
}

func (c Controller) SetModel (m *Model) {
	c.model = *m
}

func (c Controller) SetView (v *View) {
	c.view = *v
}

func (c Controller) HandleWebTraffic() {
	http.HandleFunc("/view/", c.makeHandler(c.viewHandler))
	http.HandleFunc("/edit/", c.makeHandler(c.editHandler))
	http.HandleFunc("/save/", c.makeHandler(c.saveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (c Controller) makeHandler (fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func (c Controller) viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	err := LoadPage(&c.model, title)
	if err != nil {
		http.Redirect(w, r, "/edit/" + title, http.StatusFound)
		return
	}
	c.view.RenderTemplate(w, "view", &c.model)
}

func (c Controller) editHandler(w http.ResponseWriter, r *http.Request, title string) {
	err := LoadPage(&c.model, title)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	c.view.RenderTemplate(w, "edit", &c.model)
}

func (c Controller) saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	err := Save(&c.model, title, []byte(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/" + title, http.StatusFound)
}