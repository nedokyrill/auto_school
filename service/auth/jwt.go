package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

func CreateJWT(secret []byte, userID int) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}
	JWTExpiration, _ := strconv.Atoi(os.Getenv("JWT_EXPIRATION"))
	expiration := time.Second * time.Duration(JWTExpiration)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(userID),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})

	tokenStr, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
