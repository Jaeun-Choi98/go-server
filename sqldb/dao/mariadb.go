package dao

import (
	"database/sql"
	"log"
	"root/model"

	_ "github.com/go-mysql-org/go-mysql/driver"
)

type MariaDB struct {
	db *sql.DB
}

func NewMariaDB(dbCon string) DaoTestInterfcae {
	db, err := sql.Open("mysql", dbCon)
	if err != nil {
		log.Println("failed to connect db")
		return nil
	}
	return &MariaDB{
		db: db,
	}
}

func (m *MariaDB) Ping() error {
	return m.db.Ping()
}

func (m *MariaDB) Close() {
	m.db.Close()
}

func (m *MariaDB) InsertUser(user model.User) error {
	query := `insert into test.users(user_name, age) values(?,?);`
	_, err := m.db.Exec(query, user.UserName, user.Age)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (m *MariaDB) SelectUser() ([]model.User, error) {
	query := `select * from test.users;`
	rows, err := m.db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.UserID, &user.UserName, &user.Age)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
