package services

import (
	"github.com/develop-microservices-in-go/bookstore_users-api/domain/users"
	"github.com/develop-microservices-in-go/bookstore_users-api/utils/crypto_utils"
	"github.com/develop-microservices-in-go/bookstore_users-api/utils/errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct {
}

// We need this interface so it let us mock the methods
type usersServiceInterface interface {
	GetUser(userID int64) (*users.User, *errors.RestErr)
	CreateUser(user users.User) (*users.User, *errors.RestErr)
	UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr)
	DeleteUser(userID int64) *errors.RestErr
	Search(status string) (users.Users, *errors.RestErr)
}

func (s *usersService) GetUser(userID int64) (*users.User, *errors.RestErr) {
	if userID <= 0 {
		return nil, errors.NewBadRequestError("invalid user id")
	}
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := s.GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *usersService) DeleteUser(userID int64) *errors.RestErr {
	user := &users.User{ID: userID}
	return user.Delete()
}

func (s *usersService) Search(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
