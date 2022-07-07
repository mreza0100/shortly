package services

import (
	"context"

	"github.com/mreza0100/shortly/internal/models"
	"github.com/mreza0100/shortly/internal/ports/driven"
	"github.com/mreza0100/shortly/internal/ports/services"
	er "github.com/mreza0100/shortly/pkg/errors"
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

func (s *user) Signup(ctx context.Context, email, password string) error {
	if user, err := s.cassandraRead.GetUserByEmail(ctx, email); err != nil {
		return err
	} else if user != nil && err == nil {
		return er.UsernameAlreadyExists
	}

	hashpass, err := s.passwordHasher.Hash(password)
	if err != nil {
		return err
	}

	return s.cassandraWrite.UserSignup(ctx, &models.User{
		Email:    email,
		Password: hashpass,
	})
}

func (s *user) Signin(ctx context.Context, email, password string) (string, error) {
	user, err := s.cassandraRead.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", er.InvalidEmailOrPassword
	}

	err = s.passwordHasher.Compare(user.Password, password)
	if err != nil {
		return "", er.InvalidEmailOrPassword
	}

	return s.jwtUtil.CreateToken(user.Email)
}
