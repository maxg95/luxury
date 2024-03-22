package main

import (
	"net/http"

	"luxury.maxg95/assets"

	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {
	mux := mux.NewRouter()
	mux.NotFoundHandler = http.HandlerFunc(app.notFound)

	mux.Use(app.recoverPanic)
	mux.Use(app.securityHeaders)

	fileServer := http.FileServer(http.FS(assets.EmbeddedFiles))
	mux.PathPrefix("/static/").Handler(fileServer)

	mux.HandleFunc("/", app.home).Methods("GET")
	mux.HandleFunc("/", app.insertRequest).Methods("POST")

	return mux
}
