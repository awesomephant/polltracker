package main

import (
	"encoding/json"
	"log"
	"net/http"

	"maxkohler.com/polltracker/pkg/models"
	"maxkohler.com/polltracker/pkg/usecases"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func handleListPollsters(app *models.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(405)
			w.Write([]byte("Method not allowed"))
			return
		}
		log.Println("Listing pollsters...")
		data, err := usecases.ListPollsters(10, app)
		if err != nil {
			w.Write([]byte("Could not list pollsters"))
		}
		json, err := json.Marshal(data)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(json))
	}
}

func handleAddPollster(app *models.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Allowed", http.MethodPost)
			http.Error(w, "Method not allowed.", 405)
			return
		}
		log.Println("Let's add a pollster...")
		var t models.Pollster
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&t)
		log.Println(t)
		usecases.AddPollster(t, app)
	}
}

func handleDeletePollster(app *models.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			w.Header().Set("Allow", http.MethodGet)
			http.Error(w, "Method not allowed.", 405)
			return
		}

		log.Println("Let's delete a pollster...")
		var t models.DeletePollsterTransaction
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&t)
		usecases.DeletePollster(t, app)
	}
}

func handleUpdatePollster(app *models.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Let's delete a pollster...")
		var t models.UpdatePollsterTransaction
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&t)
		usecases.UpdatePollster(t, app)
	}
}
