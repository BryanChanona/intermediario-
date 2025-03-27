package main

import (
	"fmt"
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
	fmt.Println("Servidor WebSocket iniciado en :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error iniciando servidor:", err)
	}
}
