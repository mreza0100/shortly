package services

import (
	"context"
	"log"
	"net/mail"
	"os"

	"github.com/mreza0100/shortly/internal/models"
	"github.com/mreza0100/shortly/internal/pkg/customerror"
	"github.com/mreza0100/shortly/internal/pkg/jwt"
	"github.com/mreza0100/shortly/internal/ports"
	password_hasher "github.com/mreza0100/shortly/pkg/password"
)

// User Service Dependencies
type UserServiceDep struct {
	StorageRead    ports.StorageReadPort
	StorageWrite   ports.StorageWritePort
	PasswordHasher password_hasher.PasswordHasher
	JwtUtils       jwt.JWTHelper
}

// Constructor of user service
func NewUserService(opt *UserServiceDep) ports.UserServicePort {
	return &user{
		storageRead:    opt.StorageRead,
		storageWrite:   opt.StorageWrite,
		passwordHasher: opt.PasswordHasher,
		jwtUtils:       opt.JwtUtils,
		errLogger:      log.New(os.Stderr, "UserService: ", log.LstdFlags),
	}
}

// User service implementation
type user struct {
	storageRead    ports.StorageReadPort
	storageWrite   ports.StorageWritePort
	passwordHasher password_hasher.PasswordHasher
	jwtUtils       jwt.JWTHelper
	errLogger      *log.Logger
}

// Service Signup Register new user
func (s *user) Signup(ctx context.Context, email, password string) error {
	// Check if email is valid
	if _, err := mail.ParseAddress(email); err != nil {
		return customerror.InvalidEmail
	}

	// Check if user already exists
	_, err := s.storageRead.GetUserByEmail(ctx, email)
	if err == nil {
		// If error is nil, it means that storage successfully founded user
		// So a user with this email already exists
		return customerror.EmailAlreadyExists
	} else if err != customerror.NotFound {
		// If error is not, Not Found, it means that storage failed to find user
		s.errLogger.Printf("Error getting user by email: %v", err)
		return customerror.GeneralFailure
	}

	// Hash the password
	hashedPassword, err := s.passwordHasher.Hash(password)
	if err != nil {
		s.errLogger.Printf("Error hashing password: %v", err)
		return customerror.GeneralFailure
	}

	// Save new user to storage
	return s.storageWrite.SaveUser(ctx, &models.User{
		Email:    email,
		Password: hashedPassword,
	})
}

// Service Signin new user. return JWT token
func (s *user) Signin(ctx context.Context, email, password string) (string, error) {
	// Search for user by email
	user, err := s.storageRead.GetUserByEmail(ctx, email)
	if err != nil {
		// If error is Not Found it means that user with this email not exists
		// And since we should not tell user that email not exists
		// We just return Invalid Credentials error which means ether email or password is wrong
		if err == customerror.NotFound {
			return "", customerror.InvalidCredentials
		}
		s.errLogger.Printf("Error getting user by email: %v", err)
		return "", customerror.GeneralFailure
	}
	// Just another check to make sure that we have the user
	if user == nil {
		return "", customerror.InvalidCredentials
	}

	// Check if password is correct
	if err := s.passwordHasher.Compare(user.Password, password); err != nil {
		// If error is not nil, it means that password is wrong
		return "", customerror.InvalidCredentials
	}

	// User is valid, generate JWT token
	token, err := s.jwtUtils.CreateToken(user.Id)
	if err != nil {
		s.errLogger.Printf("Error creating token: %v", err)
		return "", customerror.GeneralFailure
	}
	// Return token
	return token, nil
}
