package user

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"newWebServer/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (store *Store) GetUserByEmail(email string) (*types.User, error) {
	user := new(types.User)
	err := store.db.QueryRow("SELECT * FROM users WHERE email = ?", email).Scan(
		&user.ID,
		&user.FirstName,
		&user.Nickname,
		&user.Email)
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (store *Store) CreateUser(user types.User) error {
	return nil
}
