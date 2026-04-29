package user

import (
	"github.com/noyandey/go-ddd-starter/domain"
	userHandler "github.com/noyandey/go-ddd-starter/rest/handlers/user"
)

type Service interface {
	userHandler.Service
}

type UserRepo interface {
	Create(user domain.User) (*domain.User, error)
	Find(email string, password string) (*domain.User, error)
}
