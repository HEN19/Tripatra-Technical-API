package UserService

import "github.com/api-skeleton/model"

type UserService struct {
	Users map[string]*model.User // Simulated database
}

func NewUserService() *UserService {
	return &UserService{Users: make(map[string]*model.User)}
}
