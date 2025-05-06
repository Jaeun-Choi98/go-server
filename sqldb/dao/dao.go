package dao

import "root/model"

type DaoInterface interface {
	Close()
}

type DaoTestInterfcae interface {
	Ping() error
	InsertUser(user model.User) error
	SelectUser() ([]model.User, error)
	Close()
}
