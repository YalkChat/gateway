package encryption

import "golang.org/x/crypto/bcrypt"

type BcryptService struct{}

func (b *BcryptService) Encrypt(data string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (b *BcryptService) Decrypt(data string, hashedData string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedData), []byte(data))
}
