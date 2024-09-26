package types

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(User) error
}

type RegisterUserPayload struct {
	FirstName string `json:"first_name"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Nickname  string    `json:"nickname"`
	Level     string    `json:"level"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	SteamData SteamData `json:"steam_data"`
}

type SteamData struct {
	UserId             string `json:"user_id"`
	SteamFirstName     string `json:"steam_first_name"`
	SteamNickname      string `json:"steam_nickname"`
	SteamGamesQuantity string `json:"steam_games_quantity"`
}
