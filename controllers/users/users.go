package users

import (
	"net/http"
	"strconv"

	"github.com/develop-microservices-in-go/bookstore_users-api/domain/users"
	"github.com/develop-microservices-in-go/bookstore_users-api/services"
	"github.com/develop-microservices-in-go/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func getUserID(userIDParam string) (int64, *errors.RestErr) {
	userID, err := strconv.ParseInt(userIDParam, 10, 64)
	if err != nil {
		return 0, errors.NewBadRequestError("user id should be a number")
	}
	return userID, nil
}

func Create(ctx *gin.Context) {
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

func Get(ctx *gin.Context) {
	userID, userErr := getUserID(ctx.Param("user_id"))
	if userErr != nil {
		ctx.JSON(userErr.Status, userErr)
		return
	}
	user, getErr := services.GetUser(userID)
	if getErr != nil {
		ctx.JSON(getErr.Status, getErr)
	}
	ctx.JSON(http.StatusOK, user)
}

func Search(ctx *gin.Context) {
	ctx.String(http.StatusNotImplemented, "implement me")
}

func Update(ctx *gin.Context) {
	// Get userID query param
	userID, userErr := getUserID(ctx.Param("user_id"))
	if userErr != nil {
		ctx.JSON(userErr.Status, userErr.Error)
		return
	}
	// Get JSON body
	var user users.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status, restErr)
		return
	}

	user.ID = userID

	isPartial := ctx.Request.Method == http.MethodPatch

	result, err := services.UpdateUser(isPartial, user)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func Delete(ctx *gin.Context) {
	userID, userErr := getUserID(ctx.Param("user_id"))
	if userErr != nil {
		ctx.JSON(userErr.Status, userErr)
		return
	}

	if err := services.DeleteUser(userID); err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
