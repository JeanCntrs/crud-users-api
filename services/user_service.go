package services

import (
	"github.com/JeanCntrs/api-intermediate-level/entities"
	"github.com/JeanCntrs/api-intermediate-level/repository"
)

type UserService interface {
	FindAll() []entities.User
	FindOne(id string) (entities.User, error)
	Create(entities.User) entities.User
	Update(id string, user entities.User) entities.User
	Delete(id string) string
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService() UserService {
	return &userService{repository: repository.NewUserRepository()}
}

func (us *userService) FindAll() []entities.User {
	return us.repository.FindAll()
}

func (us *userService) FindOne(id string) (entities.User, error) {
	return us.repository.FindOne(id)
}

func (us *userService) Create(user entities.User) entities.User {
	return us.repository.Create(user)
}

func (us *userService) Update(id string, user entities.User) entities.User {
	return us.repository.Update(id, user)
}

func (us *userService) Delete(id string) string {
	return us.repository.Delete(id)
}
