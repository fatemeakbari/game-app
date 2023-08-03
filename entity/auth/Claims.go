package auth

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	jwt.RegisteredClaims
	UserID uint
	Role   string
}

func (c Claims) Valid() error {
	return c.RegisteredClaims.Valid()
}
