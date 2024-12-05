package drivers

import (
	"github.com/eatonphil/gosqlite"
	"maxkohler.com/polltracker/pkg/models"
)

type SQLitePollRepository struct {
	DB *gosqlite.Conn
}

func (repo *SQLitePollRepository) Migrate() error {
	stmt := `CREATE TABLE IF NOT EXISTS polls (
				id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
				pollster_id INTEGER,
				created DATETIME DEFAULT CURRENT_TIMESTAMP,
				FOREIGN KEY(pollster_id) REFERENCES pollsters(id)
			);
		CREATE TABLE IF NOT EXISTS poll_questions (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			poll_id INTEGER,
			questionText STRING,
			FOREIGN KEY(poll_id) REFERENCES polls(id)
		);
		CREATE TABLE IF NOT EXISTS poll_results (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			poll_question_id INTEGER,
			label STRING,
			value FLOAT32,
			FOREIGN KEY(poll_question_id) REFERENCES poll_questions(id)
		);`

	err := repo.DB.Exec(stmt)
	if err != nil {
		return err
	}
	return nil
}

func (repo *SQLitePollRepository) InsertOne(np models.Poll) (int, error) {
	return 0, nil
}
func (repo *SQLitePollRepository) GetOne(id string) (models.Poll, error) {
	return models.Poll{}, nil
}
func (repo *SQLitePollRepository) GetMany(count int) ([]models.Poll, error) {
	return nil, nil
}
