package authservice

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"messagingapp/cfg"
	"messagingapp/entity"
	"time"
)

const (
	AccessTokenSubject  = "at"
	RefreshTokenSubject = "rt"
)

type JwtTokenParser interface {
	Parse(token string) (Claims, error)
}

type JwtTokenGenerator interface {
	GenerateAccessToken(user entity.User) (string, error)
	GenerateRefreshToken(user entity.User) (string, error)
}

type Service struct {
	secretKey               string
	tokenExpirationDuration time.Duration
	tokenRefreshDuration    time.Duration
}

func New(secretKey string, expirationDuration, refreshDuration time.Duration) Service {
	return Service{
		secretKey:               secretKey,
		tokenExpirationDuration: expirationDuration,
		tokenRefreshDuration:    refreshDuration,
	}
}

func (s Service) GenerateAccessToken(user entity.User) (string, error) {
	return generate(user, AccessTokenSubject, s.tokenExpirationDuration)
}

func (s Service) GenerateRefreshToken(user entity.User) (string, error) {
	return generate(user, RefreshTokenSubject, s.tokenRefreshDuration)
}

func generate(user entity.User, subject string, expireDate time.Duration) (string, error) {

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   subject,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDate)),
		},
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenStr, err := token.SignedString([]byte(cfg.TokenSecretKey))
	if err != nil {
		return "", fmt.Errorf("error in signed token, %w", err)
	}

	return tokenStr, err
}

func (s Service) Parse(tokenStr string) (Claims, error) {

	var claims Claims

	_, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(cfg.TokenSecretKey), nil
	})

	if err != nil {
		return Claims{}, errors.New("error in parsing token")
	}

	return claims, nil
}
