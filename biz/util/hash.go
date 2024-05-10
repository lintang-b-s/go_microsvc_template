package util

import (
	"lintang/go_hertz_template/biz/domain"

	"github.com/alexedwards/argon2id"
)

// HashPassword return bcrypt hashed password
func HashPassword(password string) (string, error) {
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// if err != nil {
	// 	return "", fmt.Errorf("failed to hash password: %w", err)
	// }

	// return string(hashedPassword), nil
	hash, err := argon2id.CreateHash("pa$$word", argon2id.DefaultParams)
	if err != nil {
		return "", domain.WrapErrorf(err, domain.ErrInternalServerError, domain.MessageInternalServerError)
	}
	return hash, nil
}

// CheckPassword check jika password yang diberikan cocok atau tidak dg hashed password dari database
func CheckPassword(password string, hashedPassword string) error {
	_, err := argon2id.ComparePasswordAndHash(password, hashedPassword)
	return err
	// return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
