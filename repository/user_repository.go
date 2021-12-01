package repository

import (
	"errors"
	"strconv"

	"github.com/JeanCntrs/api-intermediate-level/entities"
)

type UserRepository interface {
	FindAll() []entities.User
	FindOne(id string) (entities.User, error)
	Create(entities.User) entities.User
	Update(id string, user entities.User) entities.User
	Delete(id string) string
}

type userRepository struct {
	users []entities.User
}

func NewUserRepository() UserRepository {
	user1 := entities.User{
		Id:      0,
		Name:    "Name 1",
		Surname: "Surname 1",
		Age:     50,
		Active:  true,
	}

	user2 := entities.User{
		Id:      1,
		Name:    "Name 2",
		Surname: "Surname 2",
		Age:     40,
		Active:  true,
	}

	user3 := entities.User{
		Id:      2,
		Name:    "Name 3",
		Surname: "Surname 3",
		Age:     30,
		Active:  true,
	}

	users := []entities.User{user1, user2, user3}

	return &userRepository{users: users}
}

func (ur *userRepository) FindAll() []entities.User {
	return ur.users
}

func (ur *userRepository) FindOne(id string) (entities.User, error) {
	convId, _ := strconv.Atoi(id)

	for _, user := range ur.users {
		if convId == user.Id {
			return user, nil
		}
	}

	return entities.User{}, errors.New("user not found")
}

func (ur *userRepository) Create(user entities.User) entities.User {
	ur.users = append(ur.users, user)
	return user
}

func (ur *userRepository) Update(id string, user entities.User) entities.User {
	convId, _ := strconv.Atoi(id)
	finalUsers := []entities.User{}

	for _, originalUser := range ur.users {
		if convId == originalUser.Id {
			finalUsers = append(finalUsers, user)
		} else {
			finalUsers = append(finalUsers, originalUser)
		}
	}

	ur.users = finalUsers

	return user
}

func (ur *userRepository) Delete(id string) string {
	convId, _ := strconv.Atoi(id)

	for index, originalUser := range ur.users {
		if convId == originalUser.Id {
			ur.users = append(ur.users[:index], ur.users[index+1:]...)
			return "user deleted"
		}
	}

	return "user not found"
}
