package controller

import (
	"dainxor/we/base/logger"
	"dainxor/we/db"
	"dainxor/we/models"
	"dainxor/we/types"

	"net/http"

	"github.com/gin-gonic/gin"
)

type userType struct{}

var User userType

func (userType) Create(c *gin.Context) {
	var body models.UserCreate

	if err := c.ShouldBindJSON(&body); err != nil {
		logger.Error(err.Error())
		logger.Error("Failed to create user: JSON request body is invalid")
		logger.Error("Request body: ", c.Request.Body)
		logger.Error("Expected body: ", "{username: string, email: string}")

		c.JSON(http.StatusBadRequest,
			models.Error(
				types.Http.BadRequest(),
				"conflict",
				"Email is already in use",
				"Expected body: {username: string, email: string}",
			))
		return
	}

	logger.Debug("Creating user: ", body)
	logger.Debug("Username: ", body.Username)
	logger.Debug("NameTag: ", body.NameTag)
	logger.Debug("Email: ", body.Email)

	result := db.User.CreateUser(body)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	user := result.Value()
	c.JSON(http.StatusCreated, user.ToResponse())
}

func (userType) GetAll(c *gin.Context) {
	result := db.User.GetAllUsers()

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	users := result.Value()
	response := types.Map(users, func(u models.UserDB) models.UserResponse { return u.ToResponse() })

	c.JSON(http.StatusOK, response)
}
func (userType) GetByID(c *gin.Context) {
	result := db.User.GetUserByID(c.Param("id"))

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	user := result.Value()
	c.JSON(http.StatusOK, user.ToResponse())
}
func (userType) GetAllByStatusID(c *gin.Context) {
	result := db.User.GetAllUsersByStatusID(c.Param("id"))

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	users := result.Value()
	response := types.Map(users, func(u models.UserDB) models.UserResponse { return u.ToResponse() })

	c.JSON(http.StatusOK, response)
}
func (userType) UpdateByID(c *gin.Context) {
	var body models.UserUpdate

	if err := c.ShouldBindJSON(&body); err != nil {
		logger.Error(err.Error())
		logger.Error("Failed to update user: JSON request body is invalid")
		logger.Error("Request body: ", c.Request.Body)
		logger.Error("Expected body: ", "{username: string, email: string}")

		c.JSON(http.StatusBadRequest,
			models.Error(
				types.Http.BadRequest(),
				"bad_request",
				"JSON request body is invalid",
				"Expected body: {username: string, email: string}",
			))
		return
	}

	response := db.User.UpdateUserByID(c.Param("id"), body)

	if response.IsErr() {
		err := response.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	c.JSON(http.StatusOK, response.Value().ToResponse())
}
func (userType) DeleteByID(c *gin.Context) {
	result := db.User.DeleteUserByID(c.Param("id"))

	if result.IsErr() {
		c.JSON(http.StatusInternalServerError, result.Error())
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
