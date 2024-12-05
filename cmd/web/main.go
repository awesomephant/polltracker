package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/eatonphil/gosqlite"
	drivers "maxkohler.com/polltracker/pkg/drivers/sqlite"
	"maxkohler.com/polltracker/pkg/models"
	"maxkohler.com/polltracker/pkg/usecases"
)

func main() {
	flag.Parse()
	db, err := gosqlite.Open("../../prod.db")
	if err != nil {
		log.Println("No DB found")
	}
	defer db.Close()

	app := &models.Application{
		Port:               ":8000",
		PollsterRepository: &drivers.SQLitePollsterRepository{DB: db},
		PollRepository:     &drivers.SQLitePollRepository{DB: db},
		DB:                 db,
	}

	usecases.MigratePollsters(app)
	usecases.MigratePolls(app)

	mux := http.NewServeMux()
	go mux.HandleFunc("/", handleIndex)
	go mux.HandleFunc("/api/pollster/create", handleAddPollster(app))
	go mux.HandleFunc("/api/pollster/list", handleListPollsters(app))
	go mux.HandleFunc("/api/pollster/delete", handleDeletePollster(app))
	go mux.HandleFunc("/api/pollster/update", handleUpdatePollster(app))

	log.Println("Listening on localhost:8000")
	err = http.ListenAndServe(app.Port, mux)
	log.Fatal(err)
}
