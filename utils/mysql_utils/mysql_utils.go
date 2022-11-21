package mysql_utils

import (
	"strings"

	"github.com/develop-microservices-in-go/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	// Try to convet err to mysql.MySQLError
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		// We was not able to handle cause is not a MySQLError
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		// We use the word "database" here so we don't tell what type of databse we are using
		return errors.NewInternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case 1062: // Duplicated key
		return errors.NewBadRequestError("invalid data")
	}
	return errors.NewInternalServerError("error processing request")
}
