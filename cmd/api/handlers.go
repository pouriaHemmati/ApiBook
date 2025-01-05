package main

import (
	"ApiBook/internal/data"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	data := map[string]string{
		"status":      "available",
		"environment": app.config.evn,
		"version":     version,
	}
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	js = append(js, '\n')
	w.Header().Set("Content-Type", "application.json")
	w.Write(js)
}

func (app *application) getCreateBooksHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		books := []data.Book{
			{
				ID:        1,
				CreateAt:  time.Now(),
				Title:     "Echoes",
				Published: 2008,
				Page:      300,
				Genres:    []string{"Action", "Fiction"},
				Rating:    8.5,
				Version:   5,
			},
			{
				ID:        2,
				CreateAt:  time.Now(),
				Title:     "Alis",
				Published: 1940,
				Page:      250,
				Genres:    []string{"Action", "Fiction"},
				Rating:    7.5,
				Version:   3,
			},
		}
		if err := app.writeJson(w, http.StatusOK, books); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	if r.Method == http.MethodPost {
		fmt.Fprintln(w, "Add a new book")
	}
}

func (app *application) getUpdateAndDeleteBooksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.getBook(w, r)
	case http.MethodPut:
		app.updateBook(w, r)
	case http.MethodDelete:
		app.deleteBook(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

}

func (app *application) getBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bade request", http.StatusBadRequest)
	}

	book := data.Book{
		ID:        idInt,
		CreateAt:  time.Now(),
		Title:     "Echoes",
		Published: 2008,
		Page:      300,
		Genres:    []string{"Action", "Fiction"},
		Rating:    8.5,
		Version:   5,
	}
	if err := app.writeJson(w, http.StatusOK, book); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}

func (app *application) updateBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bade request", http.StatusBadRequest)
	}
	fmt.Fprintf(w, "Update book of id: %d", idInt)
}

func (app *application) deleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bade request", http.StatusBadRequest)
	}
	fmt.Fprintf(w, "Delete book of id: %d", idInt)
}
