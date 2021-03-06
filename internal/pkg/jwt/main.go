package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mreza0100/shortly/internal/pkg/customerror"
)

const (
	emailKey = "email"
	expKey   = "exp"
)

type JWTHelper interface {
	CreateToken(email string) (token string, err error)
	ParseToken(token string) (email string, err error)
	IsTokenValid(token string) (isValid bool)
}

func New(secret string, expire time.Duration) JWTHelper {
	return &jwtHelper{
		secret: []byte(secret),
		expire: expire,
	}
}

type jwtHelper struct {
	secret []byte
	expire time.Duration
}

func (h *jwtHelper) CreateToken(email string) (token string, err error) {
	if email == "" {
		return "", customerror.InvalidEmail
	}

	claims := jwt.MapClaims{
		emailKey: email,
		expKey:   time.Now().Add(time.Hour * h.expire).Unix(),
	}

	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return tokenWithClaims.SignedString(h.secret)
}

func (h *jwtHelper) ParseToken(token string) (email string, err error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return h.secret, nil
	})
	if err != nil {
		return "", customerror.InvalidToken
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return "", customerror.InvalidToken
	}

	exp, ok := claims[expKey].(float64)
	if !ok {
		return "", customerror.InvalidToken
	}
	if time.Unix(int64(exp), 0).Before(time.Now()) {
		return "", customerror.ExpiredToken
	}

	email, ok = claims[emailKey].(string)
	if !ok {
		return "", customerror.InvalidToken
	}
	return email, nil
}

func (h *jwtHelper) IsTokenValid(token string) (isValid bool) {
	_, err := h.ParseToken(token)
	return err == nil
}
