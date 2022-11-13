package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {

	router := httprouter.New()

	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()

		header.Set("content-type", "application/json")
	})

	router.HandlerFunc(http.MethodGet, "/api/v1/status", app.statusHandler)

	router.HandlerFunc(http.MethodPost, "/api/v1/signin", app.SignIn)

	// User Routes

	router.HandlerFunc(http.MethodPost, "/api/v1/users/create", app.CreateNewUserPost)

	router.HandlerFunc(http.MethodPost, "/api/v1/songs/create/multiple", app.UploadBulkSongs)
	router.HandlerFunc(http.MethodPost, "/api/v1/songs/create/single", app.UploadSingleSong)
	router.HandlerFunc(http.MethodGet, "/api/v1/songs/:id", app.GetSingleSong)

	return router
}
