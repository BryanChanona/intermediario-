package mqtt

import (
	"fmt"
	"log"
	"servicio/src/Wesocket/application"
	"servicio/src/Wesocket/domain"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Procesador de mensajes MQTT
var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	payloadStr := string(msg.Payload())

	// Convertir el mensaje en estructura de dominio
	message := domain.Message{
		Topic:   msg.Topic(),
		Content: payloadStr,
		Time:    time.Now(),
	}

	// Enviar el mensaje a todos los clientes WebSocket
	application.Manager.Broadcast(message)

	fmt.Printf("Mensaje recibido en [%s]: %s\n", msg.Topic(), payloadStr)
}

// Iniciar la conexión MQTT y suscribirse a tópicos
func StartMQTTClient() {
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

	topics := []string{
		"esp32.temperatura",
		"esp32.bpm",
		"esp32.bpm2",
		"esp32.spo2",
	}

	for _, topic := range topics {
		if token := client.Subscribe(topic, 1, nil); token.Wait() && token.Error() != nil {
			log.Fatal(token.Error())
		}
		fmt.Println("Suscrito al tópico:", topic)
	}
}
