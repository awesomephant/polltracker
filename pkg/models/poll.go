package models

import (
	"time"
)

type Result struct {
	Label string
	Value float32
}

type Question struct {
	QuestionText string
	Results      []Result
}

type Poll struct {
	ID         string
	Created    time.Time
	Pollster   *Pollster
	PollDate   time.Time
	SampleSize int
	Questions  []Question
}

type PollRepository interface {
	Migrate() error
	InsertOne(Poll) (int, error)
	GetMany(count int) ([]Poll, error)
	GetOne(id string) (Poll, error)
}
