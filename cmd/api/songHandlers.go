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

func (app *application) UploadBulkSongs(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(1024 * 1024 * 1024)

	form := r.Form

	fmt.Println(form)

	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
		return
	}

	var songs []models.Song

	fmt.Println(songs)

	files, err := app.UploadFiles(r, "./uploads", true)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
		return
	}

	if err != nil {
		app.logger.Println(errors.New(fmt.Sprintf("%s", err)))
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = app.writeJSON(w, http.StatusOK, files, "data")

	if err != nil {
		app.logger.Fatalln(err)
	}

}

func (app *application) UploadSingleSong(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1024 * 1024 * 1024)

	form := r.Form

	fmt.Println(form)

	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
		return
	}

	var songs []models.Song

	fmt.Println(songs)

	s := models.Song{}

	// title

	s.Title = form.Get("song_title")

	// description
	s.Description = form.Get("song_description")

	// release date

	if len(form.Get("song_release_date")) < 1 {
		s.ReleaseDate = time.Now()
	} else {
		rd := form.Get("song_release_date")
		pt, err := time.Parse("2023-01-01T00:00:00:000Z", rd)
		if err != nil {
			app.logger.Println(errors.New("error parsing string to time"))
			return
		}
		s.ReleaseDate = pt
	}

	// song length

	s.SongLength, err = strconv.Atoi(form.Get("song_length"))
	if err != nil {
		app.logger.Println(errors.New("error parsing string to int"))
		return
	}

	//	 song rating time

	if len(form.Get("song_rating")) > 0 {
		s.Rating, err = strconv.Atoi(form.Get("song_rating"))
		if err != nil {
			app.logger.Println(errors.New("error parsing string to int"))
			return
		}
	}

	//	created at time

	if len(form.Get("song_created_at")) > 0 {
		s.Rating, err = strconv.Atoi(form.Get("song_created_at"))
		if err != nil {
			app.logger.Println(errors.New("error parsing string to int"))
			return
		}
	} else {
		s.CreatedAt = time.Now()
	}

	//	updated at time

	if len(form.Get("song_updated_at")) > 0 {
		s.UpdatedAt, err = time.Parse("2022-01-01T00:00:00:000Z", form.Get("song_updated_at"))
		if err != nil {
			app.logger.Println(errors.New("error parsing string to int"))
			return
		}
	} else {
		s.UpdatedAt = time.Now()
	}

	files, err := app.UploadSingleFile(r, "./uploads/userx", true)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
		return
	}

	s.OriginalFileName = files.OriginalFilename
	s.NewFileName = files.NewFileName
	s.FileSize = files.FileSize

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = app.writeJSON(w, http.StatusOK, files, "data")

	if err != nil {
		app.logger.Fatalln(err)
	}
}

func (app *application) GetSingleSong(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	fmt.Println(params)
}

/**********
Users
*/

func (app *application) CreateNewUserPost(w http.ResponseWriter, r *http.Request) {

}

func (app *application) SignIn(w http.ResponseWriter, r *http.Request) {
	return
}
