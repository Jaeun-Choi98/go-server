package handler

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

/**
 * websocket struct
 * manage connected client and lock
 */
type WebSocketManager struct {
	clients map[*websocket.Conn]bool
	msgChan chan []byte
	RWMu    sync.RWMutex
}

func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		clients: make(map[*websocket.Conn]bool),
		msgChan: make(chan []byte, 10),
	}
}

func (ws *WebSocketManager) Close() {
	close(ws.msgChan)
}

func (ws *WebSocketManager) handleEcho(w http.ResponseWriter, r *http.Request, upgrader *websocket.Upgrader) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("failed to upgrade ws")
		return
	}

	defer func() {
		ws.RWMu.Lock()
		delete(ws.clients, conn)
		ws.RWMu.Lock()
		conn.Close()
		//log.Println("exit")
	}()

	ws.RWMu.Lock()
	ws.clients[conn] = true
	ws.RWMu.Unlock()

	go ws.broadCast()

	// loop: receive message
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("failed to read massage: %v", err)
			break
		}

		log.Printf("received massage: %s", string(msg))

		ws.msgChan <- msg
	}
}

func (ws *WebSocketManager) broadCast() {
	for {
		msg := <-ws.msgChan

		ws.RWMu.RLock()
		for client := range ws.clients {
			if err := client.WriteMessage(websocket.TextMessage, msg); err != nil {
				log.Printf("failed to broadcast: %v", err)
				client.Close()
				delete(ws.clients, client)
			}
		}
		ws.RWMu.RUnlock()
	}
}
