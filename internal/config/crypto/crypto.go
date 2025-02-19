package crypto

type Crypto interface {
	EncryptPassword(password string) (string, error)
	ComparePasswords(hashedPassword, password string) error
}
