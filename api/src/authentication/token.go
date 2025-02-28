package authentication

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CreateToken return a signed token with the user permissions
func CreateToken(userId uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["emp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte("Secret"))
}
