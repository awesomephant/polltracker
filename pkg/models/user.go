package models

type User struct {
	ID    string
	name  string
	email string
}

type userRepository interface {
	Migrate() error
	InsertOne(User) (int, error)
	GetOne(id string) (User, error)
	DeleteOne(User) error
	UpdateOne(User) (User, error)
}
