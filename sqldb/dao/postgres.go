package dao

import (
	"context"
	"log"
	"root/model"

	"github.com/jackc/pgx/v5"
)

/**
 * pgx can be possible to use database/sql interface ( e.g. ```db, err := sql.Open("pgx", dbCon)``` )
 * but, The pgx interface is faster. pgx only targets PostgreSQL. So,
 * Many PostgreSQL specific features such as LISTEN / NOTIFY and COPY are not available through the database/sql interface.
 */
type Postgres struct {
	conn *pgx.Conn
}

func NewPostgres(dbCon string) DaoTestInterfcae {
	conn, err := pgx.Connect(context.Background(), dbCon)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &Postgres{
		conn: conn,
	}
}

func (p *Postgres) Ping() error {
	return p.conn.Ping(context.Background())
}

func (p *Postgres) InsertUser(user model.User) error {
	query := `insert into public.users(user_name,age) values($1,$2)`
	_, err := p.conn.Exec(context.Background(), query, user.UserName, user.Age)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (p *Postgres) SelectUser() ([]model.User, error) {
	query := `select * from public.users`
	rows, err := p.conn.Query(context.Background(), query)
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

func (p *Postgres) Close() { p.conn.Close(context.Background()) }
