package handler

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	HandshakeTimeout: time.Second * 5,
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// returning WebSocketManager is needed for calling WebSocketManager.Close()
func NewHandler() (http.Handler, *WebSocketManager) {
	r := mux.NewRouter()
	wsManager := NewWebSocketManager()
	r.HandleFunc("/test/ws", func(w http.ResponseWriter, r *http.Request) {
		wsManager.handleEcho(w, r, &upgrader)
	})
	return r, wsManager
}
