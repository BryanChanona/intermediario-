package routes

import (
	"net/http"
	"servicio/src/Wesocket/infrastructure/websocket"
)

// Configurar rutas
func SetupRoutes() {
	http.HandleFunc("/ws", websocket.WSHandler)
}
