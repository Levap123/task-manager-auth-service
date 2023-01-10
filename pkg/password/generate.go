package pass

import "golang.org/x/crypto/bcrypt"

func GeneratePasswordHash(password string) string {
	buffer, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(buffer)
}

func ComparePassword(passwordHash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}
