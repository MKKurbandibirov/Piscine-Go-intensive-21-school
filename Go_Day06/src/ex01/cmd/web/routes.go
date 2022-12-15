package main

import "net/http"

func (app *Application) routes(cfg *Config) *http.ServeMux {
	mux := new(http.ServeMux)

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/admin", admin)
	//mux.HandleFunc("/admin/create", create)

	fileServer := http.FileServer(neuteredFileSystem{http.Dir(cfg.StaticDir)})
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
