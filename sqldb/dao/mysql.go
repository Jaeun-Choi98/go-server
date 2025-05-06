package dao

import (
	"database/sql"
	"log"
	"root/model"
)

type MysqlDB struct {
	db *sql.DB
}

func NewMysqlDB(dbCon string) DaoTestInterfcae {
	db, err := sql.Open("mysql", dbCon)
	if err != nil {
		log.Println("failed to connect db")
		return nil
	}
	return &MysqlDB{
		db: db,
	}
}

func (m *MysqlDB) Ping() error {
	return m.db.Ping()
}

func (m *MysqlDB) InsertUser(user model.User) error {
	err := m.db.Ping()
	if err != nil {
		log.Println(err)
		return err
	}
	query := `insert into test.users(user_name,age) values(?,?);`
	_, err = m.db.Exec(query, user.UserName, user.Age)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (m *MysqlDB) SelectUser() ([]model.User, error) {
	var users []model.User
	query := `select * from test.users`
	rows, err := m.db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
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

func (m *MysqlDB) Close() {
	m.db.Close()
}
