package controller

import (
	"dainxor/we/auth"
	"dainxor/we/logger"
	"dainxor/we/mail"
	"dainxor/we/models"
	"dainxor/we/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserTryRegister(c *gin.Context) {
	email := c.Param("email")

	if utils.EmailUsed(email) {
		logger.Error("Failed to create user: Email is already in use")
		c.JSON(http.StatusBadRequest,
			models.ErrorResponse{
				Type:    "bad_request",
				Message: "Email is already in use",
			},
		)
		return
	}

	// Send email to user
	if !mail.VerifyEmailAddress(c, email) {
		return
	}

	mail.SendAuthMail(c, email)

	c.JSON(http.StatusNoContent, gin.H{})
}

func UserTryLogin(c *gin.Context) {
	email := c.Param("email")

	if !utils.EmailUsed(email) {
		logger.Error("Failed to login user: Email is not registered")
		c.JSON(http.StatusBadRequest,
			models.ErrorResponse{
				Type:    "bad_request",
				Message: "Email is not registered",
			},
		)
		return
	}

	// Send email to user
	mail.SendAuthMail(c, email)

	c.JSON(http.StatusNoContent, gin.H{})
}

func UserAuth(c *gin.Context) {
	email := c.Param("email")
	tokenStr := c.Query("token")
	token, err := strconv.ParseUint(tokenStr, 10, 32)

	if err != nil {
		logger.Error("Failed to authenticate user: Token is invalid")
		c.JSON(http.StatusBadRequest,
			models.ErrorResponse{
				Type:    "bad_request",
				Message: "could not authenticate user",
				Detail:  "invalid token",
			},
		)
		return
	}

	if !auth.VerifyCode(email, uint(token)) {
		logger.Error("Failed to authenticate user")
		c.JSON(http.StatusUnauthorized,
			models.ErrorResponse{
				Type:    "unauthorized",
				Message: "could not authenticate user",
			},
		)
		return
	}

	// Delete code from database
	auth.DeleteCode(email, uint(token))
	c.JSON(http.StatusNoContent, gin.H{})
}
