package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	secret []byte
	expiry time.Duration
}

type Claims struct {
	UserID string `json:"userId"`
	Email  string `json:"email"`

	jwt.RegisteredClaims
}

func NewJWTManager(
	secret string,
	expiry time.Duration,
) *JWTManager {
	return &JWTManager{
		secret: []byte(secret),
		expiry: expiry,
	}
}

func (m *JWTManager) Generate(
	userID string,
	email string,
) (string, error) {
	now := time.Now()

	claims := Claims{
		UserID: userID,
		Email:  email,

		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(now),

			ExpiresAt: jwt.NewNumericDate(
				now.Add(m.expiry),
			),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(m.secret)
}

func (m *JWTManager) Parse(
	tokenString string,
) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, errors.New("unexpected signing method")
			}

			return m.secret, nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)

	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
