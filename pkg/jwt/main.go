package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTHelper interface {
	CreateToken(email string) (token string, err error)
	ParseToken(token string) (email string, err error)
	IsTokenValid(token string) (isValid bool)
}

func New(secret string, expire time.Duration) JWTHelper {
	return &jwtHelper{
		secret: secret,
		expire: expire,
	}
}

type jwtHelper struct {
	secret string
	expire time.Duration
}

func (h *jwtHelper) CreateToken(email string) (token string, err error) {
	claims := jwt.MapClaims{}

	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * h.expire).Unix()

	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := tokenWithClaims.SignedString([]byte(h.secret))
	return t, err
}

func (h *jwtHelper) ParseToken(token string) (email string, err error) {
	at, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(h.secret), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := at.Claims.(jwt.MapClaims)
	if ok && at.Valid {
		// check if token is expired

		if time.Unix(int64(claims["exp"].(float64)), 0).Before(time.Now()) {
			return "", errors.New("token is expired")
		}

		// if time.Unix(int64(claims["exp"].(float64)), 0).Sub(time.Now()) < 0 {
		// 	return "", errors.New("Token is expired")
		// }
		email := claims["email"].(string)

		return email, nil
	}
	return "", err
}

func (h *jwtHelper) IsTokenValid(token string) (isValid bool) {
	_, err := h.ParseToken(token)
	return err == nil
}
