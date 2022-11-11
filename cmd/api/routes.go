package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {

	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/api/v1/status", app.statusHandler)
	router.HandlerFunc(http.MethodPost, "/api/v1/songs/create", SetContentType(app.CreateSong))
	//router.HandlerFunc(http.MethodGet, "/api/v1/songs/get/:id", app.getOneSong)
	//router.HandlerFunc(http.MethodGet, "/api/v1/songs", app.getAllSongs)

	return router
}
