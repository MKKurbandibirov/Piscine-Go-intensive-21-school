package main

import (
	"net/http"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	app.render(w, r, "home.page.tmpl") // -----------------------
}

func admin(w http.ResponseWriter, r *http.Request) {

}
