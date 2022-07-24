package seed

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/mreza0100/shortly/internal/models"
	"github.com/mreza0100/shortly/internal/ports"
)

const (
	password  = "my_str@ng_passWord"
	userCount = 30
	linkCount = 15
)

func SeedDatabase(services *ports.Services) {
	s := seed{
		services: services,
		users:    make([]*models.User, 0, userCount),
	}
	rand.Seed(time.Now().UnixNano())

	s.fillUsers()
	s.fillLinks()
}

type seed struct {
	services *ports.Services
	users    []*models.User
}

func (s *seed) getRandomStr(length int) string {
	str := ""

	for i := 0; i < length; i++ {
		str += fmt.Sprintf("%c", rand.Intn(26)+65)
	}

	return str
}

func (s *seed) getRandomEmail(length int) string {
	return s.getRandomStr(length) + "@gmail.com"
}

func (s *seed) getRandomLink(length int) string {
	return s.getRandomStr(length) + ".com"
}

func (s *seed) fillUsers() {
	for i := 0; i < userCount; i++ {
		user, err := s.createUser()
		if err != nil {
			log.Fatal(err)
		}

		s.users = append(s.users, user)
	}
}

func (s *seed) createUser() (*models.User, error) {
	email := s.getRandomEmail(10)
	ctx := context.Background()

	if err := s.services.User.Signup(ctx, email, password); err != nil {
		return nil, err
	}
	return &models.User{
		Email:    email,
		Password: password,
	}, nil
}

func (s *seed) fillLinks() {
	for _, user := range s.users {
		for i := 0; i < linkCount; i++ {
			if err := s.createLink(user.Id); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func (s *seed) createLink(id string) error {
	link := s.getRandomLink(10)
	ctx := context.Background()

	log.Printf("Creating link %s for user %s", link, id)
	if _, err := s.services.Link.NewLink(ctx, link, id); err != nil {
		return err
	}
	return nil
}
