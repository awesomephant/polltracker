package usecases

import (
	"log"

	"maxkohler.com/polltracker/pkg/models"
)

func MigratePolls(app *models.Application) error {
	err := app.PollRepository.Migrate()
	if err != nil {
		log.Println(err)
	}
	return nil
}

func AddPoll(app *models.Application, newPoll models.Poll) models.Poll {
	app.PollRepository.InsertOne(newPoll)
	return newPoll
}
