package passwordhasher

import "golang.org/x/crypto/bcrypt"

type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, password string) error
}

func New(salt string) PasswordHasher {
	return &passwordHasher{salt: salt}
}

type passwordHasher struct {
	salt string
}

func (ph *passwordHasher) Hash(password string) (string, error) {
	password += ph.salt
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (ph *passwordHasher) Compare(hashedPassword, password string) error {
	password += ph.salt
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
