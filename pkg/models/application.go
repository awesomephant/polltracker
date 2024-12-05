package models

type Application struct {
	Port               string
	PollRepository     PollRepository
	PollsterRepository PollsterRepository
	DB                 any
}
