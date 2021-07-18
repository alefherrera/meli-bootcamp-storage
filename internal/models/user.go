package models

import "github.com/google/uuid"

type User struct {
	UUID       uuid.UUID `json:"uuid"`
	Firstname  string    `json:"firstname"`
	Lastname   string    `json:"lastname"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	IP         string    `json:"ip"`
	MacAddress string    `json:"macAddress"`
	Website    string    `json:"website"`
	Image      string    `json:"image"`
}
