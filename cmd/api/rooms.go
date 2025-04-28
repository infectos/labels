package main

import (
	"fmt"
	"net/http"
	"roma/labels/internal/data"
	"time"
)

func (app *application) createRoomHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new room")
}

func (app *application) showRoomHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	room := data.Room{
		ID:        id,
		CreatedAt: time.Now(),
		Name:      "test",
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"room": room}, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encontered a problem and could not process your request", http.StatusInternalServerError)
	}
}
