package authservice

import (
	"errors"
	"fmt"
	"gameapp/entity/auth"
	"gameapp/model/usermodel"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	AccessTokenSubject  = "at"
	RefreshTokenSubject = "rt"
)

type JwtTokenParser interface {
	Parse(token string) (auth.Claims, error)
}

type JwtTokenGenerator interface {
	GenerateAccessToken(user usermodel.User) (string, error)
	GenerateRefreshToken(user usermodel.User) (string, error)
}

type Service struct {
	secretKey               string
	tokenExpirationDuration time.Duration
	tokenRefreshDuration    time.Duration
}

func New(config Config) Service {
	return Service{
		secretKey:               config.TokenSecretKey,
		tokenExpirationDuration: config.TokenAccessDuration,
		tokenRefreshDuration:    config.TokenRefreshDuration,
	}
}

func (s Service) GenerateAccessToken(user usermodel.User) (string, error) {
	return s.generate(user, AccessTokenSubject, s.tokenExpirationDuration)
}

func (s Service) GenerateRefreshToken(user usermodel.User) (string, error) {
	return s.generate(user, RefreshTokenSubject, s.tokenRefreshDuration)
}

func (s Service) generate(user usermodel.User, subject string, expireDate time.Duration) (string, error) {

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &auth.Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   subject,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDate)),
		},
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenStr, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", fmt.Errorf("error in signed token, %w", err)
	}

	return tokenStr, err
}

func (s Service) Parse(tokenStr string) (auth.Claims, error) {

	var claims auth.Claims

	_, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return auth.Claims{}, errors.New("error in parsing token")
	}

	return claims, nil
}
