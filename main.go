package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://3.234.181.19:1883")
	opts.SetClientID("mi-consumidor")
	opts.SetDefaultPublishHandler(messageHandler)
	opts.SetUsername("carlos")
	opts.SetPassword("carlos")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	topic := "esp32/temperatura"
	if token := client.Subscribe(topic, 1, nil); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	fmt.Println("Suscrito al tópico:", topic)

	// Mantener el consumidor activo
	for {
		time.Sleep(1 * time.Second)
	}
}

var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	const numberTemperaturaPromedio = 37

	// Convertir el payload (que es un []byte) a un string y luego a un número
	payloadStr := string(msg.Payload())

	// Convertir el string a float64 para comparación numérica
	var temperature float64
	_, err := fmt.Sscanf(payloadStr, "%f", &temperature)
	if err != nil {
		fmt.Println("Error al convertir el mensaje a número:", err)
		return
	}

	// Solo enviar a la API si la temperatura es mayor a 40
	if temperature > numberTemperaturaPromedio {
		fmt.Printf("Mensaje recibido en [%s]: %s\n", msg.Topic(), payloadStr)

		// Convertir el mensaje en JSON y enviarlo a la API
		jsonData := fmt.Sprintf(`{"topic": "%s", "message": "%s"}`, msg.Topic(), payloadStr)
		sendToAPI(jsonData)
	}
}

func sendToAPI(data string) {
	apiURL := "http://localhost:8081/temperature" // Cambia esto con la URL de tu API

	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer([]byte(data)))
	if err != nil {
		log.Println("Error enviando datos a la API:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Mensaje enviado a la API con status:", resp.Status)
}
