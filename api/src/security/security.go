package security

import "golang.org/x/crypto/bcrypt"

// Hash receive a string and put a hash on it
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func CheckPassword(passwordWithHash, passwordString string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordWithHash), []byte(passwordString))
}
