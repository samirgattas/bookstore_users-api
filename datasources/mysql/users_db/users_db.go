package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	/*
	 * We use "_" cause we do not use the package in our code. We only want to the side efect init
	 */
	_ "github.com/go-sql-driver/mysql"
)

const (
	// A safe we so we do not expose our credentials at github
	mysqlUsersUsername = "mysql_users_username"
	mysqlUsersPassword = "mysql_users_password"
	mysqlUsersHost     = "mysql_users_host"
	mysqlUsersSchema   = "mysql_users_schema"
)

var (
	Client *sql.DB

	/*
	 * We need to declare these environment variables. In VS Code:
	 * 1. Go to 'Edit Configurations'
	 * 2. Go to 'Environment variables'
	 * 3. Declare variable name and assign the corresponding value
	 *
	 * Other way is to run the next command in the terminal:
	 * export {variable_name}={value}
	 */
	username = os.Getenv(mysqlUsersUsername)
	password = os.Getenv(mysqlUsersPassword)
	host     = os.Getenv(mysqlUsersHost)
	schema   = os.Getenv(mysqlUsersSchema)
)

// An init function is executed automatically when the package is imported
func init() {
	// receives the DB's username, password, host, schema to connect and charset configured
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema)

	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		// if an error occurn we do not want to init the application
		panic(err)
	}
	// Check if we are able to ping the database
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
