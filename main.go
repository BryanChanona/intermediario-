package main

import (
	
	"log"
	"net/http"
	"servicio/src/Wesocket/infrastructure/mqtt"
	"servicio/src/Wesocket/infrastructure/routes"
)

func main() {
	// Configurar rutas
	routes.SetupRoutes()

	// Iniciar el cliente MQTT
	go mqtt.StartMQTTClient()

	// Iniciar servidor WebSocket
	
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal("Error iniciando servidor:", err)
	}
}
