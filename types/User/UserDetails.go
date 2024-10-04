package User

import "time"

type UserDetails struct {
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	MiddleName  string    `json:"middle_name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	PhoneMobile string    `json:"phone_mobile"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
}
