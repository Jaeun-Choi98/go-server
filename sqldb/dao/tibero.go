package dao

import (
	"database/sql"
	"log"
	"root/model"

	_ "github.com/alexbrainman/odbc"
)

type Tibero struct {
	db *sql.DB
}

func NewTibero(dbCon string) (DaoTestInterfcae, error) {
	db, err := sql.Open("odbc", "DSN=tbodbc")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &Tibero{
		db: db,
	}, nil
}

func (t *Tibero) Ping() error {
	return t.db.Ping()
}
func (t *Tibero) InsertUser(user model.User) error  { return nil }
func (t *Tibero) SelectUser() ([]model.User, error) { return nil, nil }
func (t *Tibero) Close()                            { t.db.Close() }
