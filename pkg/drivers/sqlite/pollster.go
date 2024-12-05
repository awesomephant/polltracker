package drivers

import (
	"fmt"
	"log"

	"github.com/eatonphil/gosqlite"
	"maxkohler.com/polltracker/pkg/models"
)

type SQLitePollsterRepository struct {
	DB *gosqlite.Conn
}

func (repo *SQLitePollsterRepository) Migrate() error {
	stmt := `CREATE TABLE IF NOT EXISTS pollsters (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		title VARCHAR(255) NOT NULL,
		created DATETIME DEFAULT CURRENT_TIMESTAMP,
		deleted_on DATETIME
	)`
	err := repo.DB.Exec(stmt)

	if err != nil {
		return err
	}
	return nil
}

func (repo *SQLitePollsterRepository) InsertOne(np models.Pollster) (int, error) {
	stmt := `INSERT INTO pollsters (title) VALUES(?)`
	err := repo.DB.Exec(stmt, np.Title)
	id := repo.DB.LastInsertRowID()

	if err != nil {
		return 0, err
	}
	return int(id), nil
}
func (repo *SQLitePollsterRepository) DeleteOne(t models.DeletePollsterTransaction) error {
	stmt := `DELETE FROM pollsters WHERE id = ?`
	err := repo.DB.Exec(stmt, t.Id)
	return err
}

func (repo *SQLitePollsterRepository) GetOne(id string) ([]models.Pollster, error) {
	return nil, nil
}

func (repo *SQLitePollsterRepository) GetMany(count int) ([]models.Pollster, error) {
	var result []models.Pollster
	stmt, err := repo.DB.Prepare(`SELECT * FROM pollsters LIMIT ?`, count)
	if err != nil {
		log.Println("Failed")
	}
	defer stmt.Close()
	for {
		hasRow, err := stmt.Step()
		if err != nil {
			return nil, fmt.Errorf("Step failed while querying pollsters: %v", err)
		}
		if !hasRow {
			break
		}
		var title string
		var id string
		err = stmt.Scan(&id, &title)
		if err != nil {
			return nil, fmt.Errorf("Scan failed while querying pollsters: %v", err)
		}
		np := models.Pollster{ID: id, Title: title}
		result = append(result, np)
	}
	return result, nil
}

func (repo *SQLitePollsterRepository) UpdateOne(t models.UpdatePollsterTransaction) error {
	return nil
}

func (repo *SQLitePollsterRepository) ReplaceAll(np []models.Pollster) error {
	err := repo.DB.Begin()
	if err != nil {
		return err
	}

	repo.DB.Exec("DELETE FROM pollsters")
	repo.DB.Exec("UPDATE SQLITE_SEQUENCE SET SEQ=0 WHERE NAME='pollsters';")
	repo.DB.Exec("DELETE FROM SQL_SEQUENCE WHERE name='pollsters'")

	for _, p := range np {
		repo.DB.Exec("INSERT INTO pollsters (title) VALUES(?)", p.Title)
	}

	err = repo.DB.Commit()
	if err != nil {
		return err
	}
	return err
}
