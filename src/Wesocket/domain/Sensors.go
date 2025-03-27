package domain

import "time"

// Message representa un mensaje de MQTT que se enviará a WebSocket
type Message struct {
	Topic   string    `json:"topic"`
	Content string    `json:"message"`
	Time    time.Time `json:"time"`
}