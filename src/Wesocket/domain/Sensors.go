package domain


// Message representa un mensaje de MQTT que se enviará a WebSocket
type Message struct {
	IdDevice    int     `json:"id_device"`
	IdUser      int     `json:"id_user"`
	Bpm      int `json:"bpm"`
	Spo2     int `json:"spo2"`
	Bpm2     int `json:"bpm2"`
	Temperatura float64 `json:"temperatura"`

}