package services

import (
	"context"
	"net/mail"

	"github.com/mreza0100/shortly/internal/models"
	"github.com/mreza0100/shortly/internal/ports"
	er "github.com/mreza0100/shortly/pkg/errors"
	"github.com/mreza0100/shortly/pkg/jwt"
	password_hasher "github.com/mreza0100/shortly/pkg/password"
)

type UserServiceOptions struct {
	CassandraRead  ports.CassandraReadPort
	CassandraWrite ports.CassandraWritePort
	PasswordHasher password_hasher.PasswordHasher
	JwtUtils       jwt.JWTHelper
}

func NewUserService(opt UserServiceOptions) ports.UserServicePort {
	return &user{
		cassandraRead:  opt.CassandraRead,
		cassandraWrite: opt.CassandraWrite,
		passwordHasher: opt.PasswordHasher,
		jwtUtils:       opt.JwtUtils,
	}
}

type user struct {
	cassandraRead  ports.CassandraReadPort
	cassandraWrite ports.CassandraWritePort
	passwordHasher password_hasher.PasswordHasher
	jwtUtils       jwt.JWTHelper
}

func (s *user) Signup(ctx context.Context, email, password string) error {
	if _, err := mail.ParseAddress(email); err != nil {
		return er.InvalidEmail
	}

	_, err := s.cassandraRead.GetUserByEmail(ctx, email)
	if err != nil {
		if err != er.NotFound {
			return er.GeneralFailure
		}
	} else {
		return er.EmailAlreadyExists
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

	return s.jwtUtils.CreateToken(user.Email)
}
