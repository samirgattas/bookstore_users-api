package users

import (
	"fmt"

	"github.com/develop-microservices-in-go/bookstore_users-api/datasources/mysql/users_db"
	"github.com/develop-microservices-in-go/bookstore_users-api/utils/date_utils"
	"github.com/develop-microservices-in-go/bookstore_users-api/utils/errors"
	"github.com/develop-microservices-in-go/bookstore_users-api/utils/mysql_utils"
)

/*
 * DAO (Data Access Object): it is the object that it is used to interact with de DB.
 * Here, it is the logic to persist and get element from the DB. It is the access layer to the DB.
 */

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?,?,?,?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id = ?;"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	// A statement is a connection so it's necessary to close it after use it
	stmt, err := users_db.Client.Prepare(queryGetUser)
	// Check if it was an error preparin the statement
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)

	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	// A statement is a connection so it's necessary to close it after use it cause we can run out of connections
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	// Check if it was an error preparin the statement
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		return mysql_utils.ParseError(err)
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
