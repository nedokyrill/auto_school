package types

import (
	User "newWebServer/types/User"
	"time"
)

type StudentStore interface {
}

type Student struct {
	User               User.User    `json:"user"`
	Car                Car          `json:"car"`
	Transmission       Transmission `json:"transmission"`
	Address            Address      `json:"address"`
	BanEnabled         bool         `json:"ban_enabled"`
	StartEducationDate time.Time    `json:"start_education_date"`
	PassedExams        PassedExams  `json:"passed_exams"`
}
