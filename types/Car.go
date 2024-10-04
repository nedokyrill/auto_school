package types

import "newWebServer/types/User"

type Transmission int

const (
	AUTOMATIC Transmission = iota + 1
	MECHANISM
)

type Car struct {
	ID                 int          `json:"id"`
	Name               string       `json:"name"`
	RegistrationNumber string       `json:"registration_number"`
	Color              string       `json:"color"`
	Engine             string       `json:"engine"`
	Transmission       Transmission `json:"transmission"`
	Teacher            User.User    `json:"teacher"`
}
