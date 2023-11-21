package security

import "golang.org/x/crypto/bcrypt"

func GenerateHash(password string) (string, error) {
	const cost = 14
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	return string(passwordHash), err
}
