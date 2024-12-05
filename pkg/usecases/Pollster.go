package usecases

import (
	"encoding/json"
	"log"
	"os"

	"maxkohler.com/polltracker/pkg/models"
)

func MigratePollsters(app *models.Application) {
	err := app.PollsterRepository.Migrate()
	if err != nil {
		log.Println(err)
	}
}

func SetPollstersFromJSON(app *models.Application, path string) ([]models.Pollster, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var np []models.Pollster
	err = json.Unmarshal(data, &np)
	if err != nil {
		return nil, err
	}

	err = app.PollsterRepository.ReplaceAll(np)
	if err != nil {
		return nil, err
	}
	return np, nil
}

func ListPollsters(count int, app *models.Application) ([]models.Pollster, error) {
	result, err := app.PollsterRepository.GetMany(count)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil

}
func AddPollster(p models.Pollster, app *models.Application) (int, error) {
	id, err := app.PollsterRepository.InsertOne(p)
	if err != nil {
		// comma ok idiom
		return 0, err
	}
	return id, nil
}

func DeletePollster(t models.DeletePollsterTransaction, app *models.Application) error {
	err := app.PollsterRepository.DeleteOne(t)
	if err != nil {
		log.Println(err)
	}
	return err
}

func UpdatePollster(t models.UpdatePollsterTransaction, app *models.Application) error {
	err := app.PollsterRepository.UpdateOne(t)
	if err != nil {
		log.Println(err)
	}
	return err
}

func ReplaceAllPollsters(np []models.Pollster, app *models.Application) error {
	err := app.PollsterRepository.ReplaceAll(np)
	return err
}
