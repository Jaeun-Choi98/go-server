package dao

import (
	"database/sql"
	"log"
	"root/model"

	_ "github.com/godror/godror"
)

type OracleDB struct {
	db *sql.DB
}

// sql.Open("godror", `user="scott" password="tiger" connectString="dbhost:1521/orclpdb1"`)
func NewOracleDB(dbCon string) DaoTestInterfcae {
	db, err := sql.Open("godror", dbCon)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &OracleDB{
		db: db,
	}
}

func (o *OracleDB) Ping() error {
	return o.db.Ping()
}

// ORA-00911 err -> query에 세미콜론이나 백틱이 있다면 발생.
func (o *OracleDB) InsertUser(user model.User) error {
	query := `insert into users values(user_id.nextval, :1, :2)`
	_, err := o.db.Exec(query, user.UserName, user.Age)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (o *OracleDB) SelectUser() ([]model.User, error) {
	var users []model.User
	query := `select * from juchoi.users`
	rows, err := o.db.Query(query)
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

func (o *OracleDB) Close() { o.db.Close() }
