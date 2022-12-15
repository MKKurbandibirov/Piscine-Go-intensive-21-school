package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *Application) serverError(w http.ResponseWriter, e error) {
	trace := fmt.Sprintf("%s\n%s", e.Error(), debug.Stack())
	app.logger.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) wrongArg(w http.ResponseWriter, e error) {
	trace := fmt.Sprintf("%s\n%s", e.Error(), debug.Stack())
	app.logger.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}

func (app *Application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *Application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *Application) render(w http.ResponseWriter, r *http.Request, name string /*, td *templateData*/) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("Template %s doesn't exist", name))
		return
	}

	err := ts.Execute(w, nil) // ----------------------------
	if err != nil {
		app.serverError(w, err)
	}
}
