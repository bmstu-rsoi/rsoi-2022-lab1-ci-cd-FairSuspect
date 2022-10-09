package model

type Person struct {
	Id      int    `json:"id" gorm:"primary_key"`
	Name    string `json:"name" gorm:"not null"`
	Age     int    `json:"age,omitempty"`
	Address string `json:"address,omitempty"`
	Work    string `json:"work,omitempty"`
}
