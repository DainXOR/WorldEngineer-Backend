package controller

import (
	"dainxor/we/auth"
	"dainxor/we/db"
	"dainxor/we/logger"
	"dainxor/we/mail"
	"dainxor/we/models"
	"dainxor/we/types"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type authType struct{}

var Auth authType

func (authType) Register(c *gin.Context) {
	email := c.Param("email")
	print(email)

	if db.GetUserByEmail(email).IsOk() {
		logger.Error("Failed to create user: Email is already in use")

		c.JSON(http.StatusBadRequest,
			models.Error(
				types.Http.Conflict(),
				"conflict",
				"Email is already in use",
			),
		)
		return
	}

	if result := mail.VerifyEmailAddress(email); result.IsPresent() {
		result := result.Get()
		c.JSON(result.Code.Get(), result)
		return
	}

	if result := mail.SendAuthMail(email); result.IsErr() {
		result := result.Error()
		c.JSON(result.Code.Get(), result)
	} else {
		c.JSON(http.StatusNoContent, gin.H{})
	}
}

func (authType) Login(c *gin.Context) {
	email := c.Param("email")

	if db.GetUserByEmail(email).IsErr() {
		logger.Error("Failed to login user: Email is not registered")

		c.JSON(http.StatusBadRequest,
			models.Error(
				types.Http.BadRequest(),
				"bad_request",
				"Email is not registered",
			),
		)
		return
	}

	if result := mail.SendAuthMail(email); result.IsErr() {
		result := result.Error()
		c.JSON(result.Code.Get(), result)
	} else {
		c.JSON(http.StatusNoContent, gin.H{})
	}
}

func (authType) Authenticate(c *gin.Context) {
	email := c.Param("email")
	tokenStr := c.Query("token")
	_, err := strconv.ParseUint(tokenStr, 10, 32)

	if err != nil {
		logger.Error("Failed to authenticate user: Token is invalid")
		c.JSON(http.StatusBadRequest,
			models.Error(
				types.Http.BadRequest(),
				"bad_request",
				"could not authenticate user",
				"invalid token",
			),
		)
		return
	}

	if !auth.VerifyCode(email, tokenStr) {
		logger.Error("Failed to authenticate user")
		c.JSON(http.StatusUnauthorized,
			models.Error(
				types.Http.Unauthorized(),
				"unauthorized",
				"could not authenticate user",
			),
		)
		return
	}

	// Delete code from database
	auth.DeleteCode(email, tokenStr)
	c.JSON(http.StatusNoContent, gin.H{})
}
