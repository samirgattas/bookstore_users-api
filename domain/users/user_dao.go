package users

import (
	"fmt"
	"strings"

	"github.com/develop-microservices-in-go/bookstore_users-api/datasources/mysql/users_db"
	"github.com/develop-microservices-in-go/bookstore_users-api/utils/date_utils"
	"github.com/develop-microservices-in-go/bookstore_users-api/utils/errors"
)

/*
 * DAO (Data Access Object): it is the object that it is used to interact with de DB.
 * Here, it is the logic to persist and get element from the DB. It is the access layer to the DB.
 */

const (
	indexUniqueEmail = "UNIQUE_email"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?,?,?,?);"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
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
	// A statement is a connection so it's necessary to close it after use it
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	// Check if it was an error preparin the statement
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	/*
	 * Another way to do an insert is just pass the query with the values
	 * insertResult, err := stmt.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated)
	 * Fede León says that the 1st way is better because Prepare() method
	 * validates that the query is a valid query and performs better.
	 */

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when thying to save user: %s", err.Error()))
	}
	user.ID = userID
	return nil
}
