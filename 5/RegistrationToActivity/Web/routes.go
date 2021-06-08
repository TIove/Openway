package main

import "net/http"

func (app *Application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	userRoutes(app, mux)

	fileServer := http.FileServer(http.Dir("./Web/ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}

func userRoutes(app *Application, mux *http.ServeMux) {
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/login", app.login)
	mux.HandleFunc("/admin", app.admin)
}

