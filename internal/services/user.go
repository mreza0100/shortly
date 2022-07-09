package services

import (
	"context"
	"log"
	"net/mail"
	"os"

	"github.com/mreza0100/shortly/internal/models"
	er "github.com/mreza0100/shortly/internal/pkg/errors"
	"github.com/mreza0100/shortly/internal/pkg/jwt"
	"github.com/mreza0100/shortly/internal/ports"
	password_hasher "github.com/mreza0100/shortly/pkg/password"
)

type UserServiceDep struct {
	StorageRead    ports.StorageReadPort
	StorageWrite   ports.StorageWritePort
	PasswordHasher password_hasher.PasswordHasher
	JwtUtils       jwt.JWTHelper
}

func NewUserService(opt *UserServiceDep) ports.UserServicePort {
	return &user{
		storageRead:    opt.StorageRead,
		storageWrite:   opt.StorageWrite,
		passwordHasher: opt.PasswordHasher,
		jwtUtils:       opt.JwtUtils,
		errLogger:      log.New(os.Stderr, "UserService: ", log.LstdFlags),
	}
}

type user struct {
	storageRead    ports.StorageReadPort
	storageWrite   ports.StorageWritePort
	passwordHasher password_hasher.PasswordHasher
	jwtUtils       jwt.JWTHelper
	errLogger      *log.Logger
}

func (s *user) Signup(ctx context.Context, email, password string) error {
	if _, err := mail.ParseAddress(email); err != nil {
		return er.InvalidEmail
	}

	_, err := s.storageRead.GetUserByEmail(ctx, email)
	if err != nil {
		if err != er.NotFound {
			s.errLogger.Printf("Error getting user by email: %v", err)
			return er.GeneralFailure
		}
	} else {
		return er.EmailAlreadyExists
	}

	hashpass, err := s.passwordHasher.Hash(password)
	if err != nil {
		s.errLogger.Printf("Error hashing password: %v", err)
		return er.GeneralFailure
	}

	return s.storageWrite.UserSignup(ctx, &models.User{
		Email:    email,
		Password: hashpass,
	})
}

func (s *user) Signin(ctx context.Context, email, password string) (string, error) {
	user, err := s.storageRead.GetUserByEmail(ctx, email)
	if err != nil {
		if err == er.NotFound {
			return "", er.InvalidCredentials
		}
		s.errLogger.Printf("Error getting user by email: %v", err)
		return "", er.GeneralFailure
	}
	if user == nil {
		return "", er.InvalidCredentials
	}

	err = s.passwordHasher.Compare(user.Password, password)
	if err != nil {
		return "", er.InvalidCredentials
	}

	token, err := s.jwtUtils.CreateToken(user.Email)
	if err != nil {
		s.errLogger.Printf("Error creating token: %v", err)
		return "", er.GeneralFailure
	}
	return token, nil
}
