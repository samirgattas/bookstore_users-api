package users

import (
	"fmt"

	"github.com/develop-microservices-in-go/bookstore_users-api/utils/date_utils"
	"github.com/develop-microservices-in-go/bookstore_users-api/utils/errors"
)

/*
 * DAO (Data Access Object): it is the object that it is used to interact with de DB.
 * Here, it is the logic to persist and get element from the DB. It is the access layer to the DB.
 */

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := userDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}
	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	current := userDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.ID))
	}
	user.DateCreated = date_utils.GetNowString()
	userDB[user.ID] = user
	return nil
}
