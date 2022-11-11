package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elkcityhazard/lofi-lounge-redux/internal/models"
	"net/http"
)

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, wrap string) error {

	fmt.Println("writing json")
	wrapper := make(map[string]interface{})

	wrapper[wrap] = data

	js, err := json.Marshal(wrapper)

	if err != nil {
		return err
	}

	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(status)
	fmt.Println("working")
	_, err = w.Write(js)

	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return err
	}

	return nil
}

func (app *application) errorJSON(w http.ResponseWriter, err error) {
	type jsonError struct {
		Message string `json:"message"`
	}

	theError := jsonError{
		Message: err.Error(),
	}

	app.writeJSON(w, http.StatusBadRequest, theError, "error")
}

func (app *application) GenerateErrorJSON(method string, status int, errMsg string, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fired")
	// This needs to be a post request
	if r.Method != method {
		w.Header().Set("Access-Control-Allow-Methods", http.MethodPost)
		msg := errors.New(errMsg)

		errStruct := models.Error{
			Status:  status,
			Message: msg,
		}

		err := app.writeJSON(w, http.StatusMethodNotAllowed, errStruct, "error")

		if err != nil {
			http.Error(w, "error processing request", http.StatusInternalServerError)
			return
		}
	}
}
