package main

import (
	"flag"
	"log"

	"github.com/eatonphil/gosqlite"
	drivers "maxkohler.com/polltracker/pkg/drivers/sqlite"
	"maxkohler.com/polltracker/pkg/models"
	"maxkohler.com/polltracker/pkg/usecases"
)

func main() {
	flag.Parse()
	log.Print("Establishing database connection...")
	db, err := gosqlite.Open("../../prod.db")
	if err != nil {
		log.Println("Failed to establish database connection")
	}

	defer db.Close()

	app := &models.Application{
		DB:                 db,
		PollsterRepository: &drivers.SQLitePollsterRepository{DB: db},
		PollRepository:     &drivers.SQLitePollRepository{DB: db},
	}

	log.Print("Resetting pollsters from config...")
	ps, err := usecases.SetPollstersFromJSON(app, "../../config/pollsters.json")

	for p := range ps {
		log.Print(p)
	}
}
