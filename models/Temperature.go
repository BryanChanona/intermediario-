package models

type Temperature struct {
	Id int `json:"id,omitempty"`
	Id_user  int  `json:"id_user"`
	Date string `json:"date"`
	Time string `json:"time"`
	RegisteredMeasure float32 `json:"registeredMeasure"`

}
