package main

import (
	"errors"
	"fmt"
	"github.com/elkcityhazard/lofi-lounge-redux/internal/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"time"
)

func (app *application) getOneSong(w http.ResponseWriter, r *http.Request) {

	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		app.logger.Println(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	app.logger.Println("id is: ", id)

	song := models.Song{
		ID:          id,
		Title:       "Some Song",
		Description: "Some Description",
		Year:        2022,
		ReleaseDate: time.Date(2022, 10, 15, 0, 0, 0, 0, time.Local),
		SongLength:  180,
		Rating:      5,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		SongGenre:   []models.SongGenre{},
	}

	err = app.writeJSON(w, http.StatusOK, song, "song")

}

func (app *application) getAllSongs(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}

	fmt.Println(user)
}

func (app *application) CreateSong(w http.ResponseWriter, r *http.Request, next http.Handler) {
	// This needs to be a post request
	app.GenerateErrorJSON(http.MethodPost, http.StatusMethodNotAllowed, "method not allowed", w, r)
	next.ServeHTTP(w, r)
}

/**********
Users
*/

func (app *application) CreateNewUserPost() {

}
