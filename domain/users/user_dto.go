package users

import (
	"strings"

	"github.com/develop-microservices-in-go/bookstore_users-api/utils/errors"
)

/*
 * DTO (Data Transfer Object): It is the object that we persist in the DB.
 */
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string
}

// With this method, the same user knows how to validate by himself
func (user *User) Validate() *errors.RestErr {
	// Remove white spaces and then, to lower case
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}
