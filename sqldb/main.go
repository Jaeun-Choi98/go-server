package main

import (
	"net/http"
	"os"
	"root/handler"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	mux := mux.NewRouter()
	h := handler.NewHandler(os.Getenv("DB_CON"))
	defer h.Close()
	mux.HandleFunc("/get-vmsservinfo", h.GetVmsServinfo)
	mux.HandleFunc("/get-grp-admin", h.GetGroupAndAdmin)
	http.ListenAndServe(":8080", mux)
}
