package application

import (
	"encoding/json"
	"log"
	"servicio/src/Wesocket/domain"
	"sync"

	"github.com/gorilla/websocket"
)

// ClientManager gestiona las conexiones WebSocket activas
type ClientManager struct {
	clients map[*websocket.Conn]bool
	mu      sync.Mutex
}

var Manager = ClientManager{
	clients: make(map[*websocket.Conn]bool),
}

// Agregar un nuevo cliente WebSocket
func (cm *ClientManager) AddClient(conn *websocket.Conn) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.clients[conn] = true
}

// Remover un cliente WebSocket
func (cm *ClientManager) RemoveClient(conn *websocket.Conn) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	delete(cm.clients, conn)
}

// Enviar mensaje a todos los clientes WebSocket
func (cm *ClientManager) Broadcast(message domain.Message) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	jsonMsg, err := json.Marshal(message)
	if err != nil {
		log.Println("Error serializando mensaje:", err)
		return
	}

	for client := range cm.clients {
		err := client.WriteMessage(websocket.TextMessage, jsonMsg)
		if err != nil {
			log.Println("Error enviando mensaje a cliente:", err)
			client.Close()
			delete(cm.clients, client)
		}
	}
}
