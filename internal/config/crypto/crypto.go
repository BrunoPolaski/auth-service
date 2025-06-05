package crypto

import "github.com/BrunoPolaski/go-rest-err/rest_err"

type Crypto interface {
	EncryptPassword(password string) (string, *rest_err.RestErr)
	ComparePasswords(hashedPassword, password string) *rest_err.RestErr
}
