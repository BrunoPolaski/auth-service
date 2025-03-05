package jwt

import (
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"github.com/golang-jwt/jwt/v5"
)

type JWT interface {
	GenerateToken(claims jwt.MapClaims) (string, *rest_err.RestErr)
	ParseToken(token string) (interface{}, *rest_err.RestErr)
}
