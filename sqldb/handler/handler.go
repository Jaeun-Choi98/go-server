package handler

import (
	"root/dao"
)

type HandlerInterface interface {
	Close()
}

type myHandler struct {
	dao dao.DaoInterface
}
