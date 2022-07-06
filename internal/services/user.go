package services

import (
	"errors"

	"github.com/mreza0100/shortly/internal/entities"
	"github.com/mreza0100/shortly/internal/ports/driven"
	"github.com/mreza0100/shortly/internal/ports/services"
	"github.com/mreza0100/shortly/pkg/jwt"
	passwordhasher "github.com/mreza0100/shortly/pkg/password"
)

type UserServiceOptions struct {
	CassandraRead  driven.CassandraReadPort
	CassandraWrite driven.CassandraWritePort
	PasswordHasher passwordhasher.PasswordHasher
	JwtUtil        jwt.JWTHelper
}

func NewUserService(opt UserServiceOptions) services.UserServicePort {
	return &user{
		cassandraRead:  opt.CassandraRead,
		cassandraWrite: opt.CassandraWrite,
		passwordHasher: opt.PasswordHasher,
		jwtUtil:        opt.JwtUtil,
	}
}

type user struct {
	cassandraRead  driven.CassandraReadPort
	cassandraWrite driven.CassandraWritePort
	passwordHasher passwordhasher.PasswordHasher
	jwtUtil        jwt.JWTHelper
}

func (s *user) Signup(email, password string) error {
	if user, err := s.cassandraRead.GetUserByEmail(email); err != nil {
		return err
	} else if user != nil && err == nil {
		return errors.New("Username already exists")
	}

	hashpass, err := s.passwordHasher.Hash(password)
	if err != nil {
		return err
	}

	return s.cassandraWrite.UserSignup(&entities.User{
		Email:    email,
		Password: hashpass,
	})
}

func (s *user) Login(email, password string) (string, error) {
	user, err := s.cassandraRead.GetUserByEmail(email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("Invalid email or password")
	}

	err = s.passwordHasher.Compare(user.Password, password)
	if err != nil {
		return "", errors.New("Invalid email or password")
	}

	return s.jwtUtil.CreateToken(user.Email)
}
