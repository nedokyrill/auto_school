package user

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"newWebServer/types/User"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (store *Store) GetUserByEmail(email string) (*User.User, error) {
	user := new(User.User)
	err := store.db.QueryRow("SELECT id, firstname, email, password FROM users WHERE email = $1", email).Scan(
		&user.ID,
		&user.UserDetails.FirstName,
		&user.UserDetails.Email,
		&user.UserDetails.Password,
	)

	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (store *Store) GetUserById(id int) (*User.User, error) {
	user := new(User.User)
	err := store.db.QueryRow("SELECT id, firstname, email FROM users WHERE id = $1", id).Scan(
		&user.ID,
		&user.UserDetails.FirstName,
		&user.UserDetails.Email,
	)

	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (store *Store) CreateUser(user User.User) error {
	_, err := store.db.Exec(
		"INSERT INTO users (firstname, lastname, email, password) VALUES ($1, $2, $3, $4)",
		user.UserDetails.FirstName,
		user.UserDetails.LastName,
		user.UserDetails.Email,
		user.UserDetails.Password,
	)

	if err != nil {
		return err
	}
	return nil
}

//func (store *Store) UpdateUser(user types.User) error {
//
//	_, err := store.db.Exec(
//		"UPDATE users SET firstname = $1, lastname = $2,  password = $3 WHERE email = $4",
//		user.UserDetails.FirstName,
//		user.UserDetails.LastName,
//		user.UserDetails.Password,
//		user.UserDetails.Email,
//	)
//
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
