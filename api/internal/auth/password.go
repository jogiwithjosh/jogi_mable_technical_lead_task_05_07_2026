package auth

import "golang.org/x/crypto/bcrypt"

const (
	// DefaultCost is a good balance between security and performance.
	DefaultCost = bcrypt.DefaultCost
)

// Hash hashes a plain text password using bcrypt.
func Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		DefaultCost,
	)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// Verify compares a plain text password with a bcrypt hash.
func Verify(
	password string,
	hash string,
) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)
}
