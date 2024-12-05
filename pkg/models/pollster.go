package models

type Pollster struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Website string `json:"website"`
}

type PollsterRepository interface {
	InsertOne(Pollster) (int, error)
	GetOne(id string) ([]Pollster, error)
	GetMany(count int) ([]Pollster, error)
	DeleteOne(t DeletePollsterTransaction) error
	UpdateOne(t UpdatePollsterTransaction) error
	ReplaceAll(newPollsters []Pollster) error
	Migrate() error
}

type DeletePollsterTransaction struct {
	Id string
}

type UpdatePollsterTransaction struct {
	Id string
	P  Pollster
}
