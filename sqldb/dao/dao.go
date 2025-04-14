package dao

import "root/model"

type DaoInterface interface {
	GetGroupAndAdmin() []model.GroupAndAdmin
	GetVMSWithServeInfo() []model.VMSWithServinfo

	Close()
}

type DaoTestInterfcae interface {
	InsertUser(user model.User) error
	SelectUser() ([]model.User, error)
	Close()
}
