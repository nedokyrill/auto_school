package User

import "time"

type RoleType int

const (
	Admin RoleType = iota + 1
	Teacher
	Student
)

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
	CreateUser(User) error
	//UpdateUser(User) error
}

type RegisterUserPayload struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name"  validate:"required"`
	Email     string `json:"email"      validate:"required,email"`
	Password  string `json:"password"   validate:"required,min=5,max=20"`
}

type User struct {
	ID           int         `json:"id"`
	FullName     string      `json:"full_name"`
	UserDetails  UserDetails `json:"user_details"`
	Role         RoleType    `json:"role"`
	CreatedAt    time.Time   `json:"createdAt"`
	UpdatedAt    time.Time   `json:"updatedAt"`
	EmailEnabled bool        `json:"email_enabled"`
}

type LoginUserPayload struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
