package jwt

import (
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"github.com/golang-jwt/jwt/v5"
)

type jwtAdapter struct{}

func NewJWTAdapter() JWT {
	return &jwtAdapter{}
}

func (ja *jwtAdapter) GenerateToken(claims jwt.MapClaims) (string, *rest_err.RestErr) {
	return "", nil
}

func (ja *jwtAdapter) ParseToken(token string) (interface{}, *rest_err.RestErr) {
	return "", nil
}
