package handler

import (
	"encoding/json"
	"net/http"
	"root/dao"
)

type HandlerInterface interface {
	Close()
	GetVmsServinfo(w http.ResponseWriter, r *http.Request)
	GetGroupAndAdmin(w http.ResponseWriter, r *http.Request)
}

type myHandler struct {
	dao dao.DaoInterface
}

func NewHandler(dbCon string) HandlerInterface {
	return &myHandler{dao.NewMariaDB(dbCon)}
}

func (mh *myHandler) GetGroupAndAdmin(w http.ResponseWriter, r *http.Request) {
	data := mh.dao.GetGroupAndAdmin()
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&data)
}

func (mh *myHandler) GetVmsServinfo(w http.ResponseWriter, r *http.Request) {
	data := mh.dao.GetVMSWithServeInfo()
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&data)
}

func (mh *myHandler) Close() {
	mh.dao.Close()
}
