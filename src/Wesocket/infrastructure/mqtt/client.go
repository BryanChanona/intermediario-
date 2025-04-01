package mqtt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"servicio/src/Wesocket/application"
	"servicio/src/Wesocket/domain"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Procesador de mensajes MQTT
var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	payloadStr := string(msg.Payload())
	

	var payload domain.Message

	// Decodificar el JSON del mensaje recibido
	err := json.Unmarshal([]byte(payloadStr), &payload)
	if err != nil {
		fmt.Println("Error al decodificar el mensaje JSON:", err)
		return
	}

	application.Manager.Broadcast(payload)



	
	var apiURL string
	var data map[string]interface{}

	if payload.Spo2 != 0 { // Si hay un valor de SPO2, se enviará a la ruta de oxígeno
		apiURL = "http://localhost:8081/oxygen/"
		data = map[string]interface{}{
			"id_user":  payload.IdUser,
			"registeredMeasure": payload.Spo2,
			"id_device":    payload.IdDevice,	
		}
	} else if payload.Bpm != 0 { // Si hay un valor de BPM, se enviará a la ruta de frecuencia cardíaca
		apiURL = "http://localhost:8081/heartRate/"
		data = map[string]interface{}{
			"id_user":  payload.IdUser,
			"registeredMeasure": payload.Bpm,
			"id_device":    payload.IdDevice,
		}
	}else if payload.Bpm2 != 0 {
		
		apiURL = "http://localhost:8081/heartRate/"
		data = map[string]interface{}{
			"id_user":  payload.IdUser,
			"registeredMeasure": payload.Bpm,
			"id_device":    payload.IdDevice,
		}	
	} else if payload.Temperatura != 0 { // Si hay un valor de temperatura, se enviará a la ruta de temperatura
		apiURL = "http://localhost:8081/temperature/"
		data = map[string]interface{}{
			"id_user":  payload.IdUser,
			"registeredMeasure": payload.Temperatura,
			"id_device":    payload.IdDevice,	
		}
	}
	// Si se asignó una URL, hacer la solicitud HTTP POST
	if apiURL != "" {
		// Convertir los datos a JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error al convertir los datos a JSON:", err)
			return
		}

		// Hacer la solicitud HTTP POST a la API
		resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("Error al hacer la solicitud HTTP:", err)
			return
		}
		defer resp.Body.Close()

		// Imprimir la respuesta de la API (opcional)
		fmt.Printf("Respuesta de la API: %s\n", resp.Status)
	} else {
		fmt.Println("No se encontraron datos válidos para enviar a la API.")
	}


	// Enviar el mensaje a todos los clientes WebSocket
	
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
		"esp32.datos",
		"esp32.datos2",
		
	}

	for _, topic := range topics {
		if token := client.Subscribe(topic, 1, nil); token.Wait() && token.Error() != nil {
			log.Fatal(token.Error())
		}
		fmt.Println("Suscrito al tópico:", topic)
	}
}
