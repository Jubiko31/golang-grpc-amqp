package main

import (
	"errors"
	"fmt"
	"net/http"
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPyload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &requestPyload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	user, err := app.Models.User.GetByEmail(requestPyload.Email)
	if err != nil {
		app.errorJSON(w, errors.New("Invalid credentials"), http.StatusBadRequest)
		return
	}

	valid, err := user.PasswordMatches(requestPyload.Password)
	if err != nil || !valid {
		app.errorJSON(w, errors.New("Invalid credentials"), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in successfully. User: %s", user.Email),
		Data:    user,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}
