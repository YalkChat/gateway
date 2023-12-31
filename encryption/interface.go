package encryption

type Service interface {
	Encrypt(data string) (string, error)
	Decrypt(data string, hashedData string) error
}
