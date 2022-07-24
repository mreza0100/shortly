package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mreza0100/shortly/internal/pkg/customerror"
)

const (
	idKey  = "id"
	expKey = "exp"
)

type JWTHelper interface {
	CreateToken(id string) (token string, err error)
	ParseToken(token string) (id string, err error)
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

func (h *jwtHelper) CreateToken(id string) (token string, err error) {
	claims := jwt.MapClaims{
		idKey:  id,
		expKey: time.Now().Add(time.Hour * h.expire).Unix(),
	}

	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return tokenWithClaims.SignedString(h.secret)
}

func (h *jwtHelper) ParseToken(token string) (id string, err error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) { return h.secret, nil })
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

	id, ok = claims[idKey].(string)
	if !ok {
		return "", customerror.InvalidToken
	}
	return id, nil
}

func (h *jwtHelper) IsTokenValid(token string) (isValid bool) {
	_, err := h.ParseToken(token)
	return err == nil
}
