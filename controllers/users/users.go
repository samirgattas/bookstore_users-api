package users

import (
	"net/http"
	"strconv"

	"github.com/develop-microservices-in-go/bookstore_users-api/domain/users"
	"github.com/develop-microservices-in-go/bookstore_users-api/services"
	"github.com/develop-microservices-in-go/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var user users.User
	// WAY 1
	/*
		bytes, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			// TODO: Handler error
			return
		}
		if err := json.Unmarshal(bytes, &user); err != nil {
			// TODO: Handle json error
			return
		}
	*/

	// WAY 2
	if err := ctx.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		ctx.JSON(saveErr.Status, saveErr)
		return
	}
	ctx.JSON(http.StatusCreated, result)
}

func GetUser(ctx *gin.Context) {
	userID, userErr := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		ctx.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userID)
	if getErr != nil {
		ctx.JSON(getErr.Status, getErr)
	}
	ctx.JSON(http.StatusOK, user)
}

func SearchUser(ctx *gin.Context) {
	ctx.String(http.StatusNotImplemented, "implement me")
}
