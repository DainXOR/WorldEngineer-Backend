package controller

import (
	"dainxor/we/base/logger"
	"dainxor/we/db"
	"dainxor/we/models"
	"dainxor/we/types"
	"dainxor/we/utils"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type authType struct{}

var Auth authType

func (authType) Register(c *gin.Context) {
	email := c.Param("email")
	print(email)

	if db.User.GetUserByEmail(email).IsOk() {
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

	if db.Auth.GetValidCodeByEmail(email).IsOk() {
		c.JSON(types.Http.Conflict().Get(),
			models.Error(
				types.Http.Conflict(),
				"conflict",
				"Email is already in use",
				"Check your email for the verification code, or try again in a few minutes. Also try looking in your spam folder",
			),
		)
		return
	}

	if err := db.Mail.VerifyEmailAddress(email); err.IsPresent() {
		err := err.Get()
		c.JSON(err.Code.Get(), err)
		return
	}

	if result := db.Mail.SendAuthMail(email); result.IsErr() {
		result := result.Error()
		c.JSON(result.Code.Get(), result)
	} else {
		c.JSON(http.StatusNoContent, gin.H{})
	}
}

func (authType) Login(c *gin.Context) {
	email := c.Param("email")

	if db.User.GetUserByEmail(email).IsErr() {
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

	if result := db.Mail.SendAuthMail(email); result.IsErr() {
		result := result.Error()
		c.JSON(result.Code.Get(), result)
	} else {
		c.JSON(http.StatusNoContent, gin.H{})
	}
}

func (authType) Verify(c *gin.Context) {
	email := c.Param("email")
	tokenStr := c.Query("token")
	_, err := strconv.ParseUint(tokenStr, 10, 32)

	logger.Info("Verifying user")
	logger.Debug("Email: " + email)
	logger.Debug("Token: " + tokenStr)

	if err != nil {
		logger.Warning("Failed to authenticate user: Token is invalid")
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

	result := db.Auth.ConsumeCodeByEmail(email, tokenStr)

	if result.IsErr() {
		logger.Warning("Failed to authenticate user")
		logger.Info("Reason: " + result.Error().Detail)

		c.JSON(http.StatusUnauthorized,
			models.Error(
				types.Http.Unauthorized(),
				"unauthorized",
				"could not authenticate user",
				result.Error().Message,
			),
		)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func (authType) CreateAccount(c *gin.Context) {
	var body models.UserCreate

	if err := c.ShouldBindJSON(&body); err != nil {
		logger.Error(err.Error())
		logger.Error("Failed to create user: JSON request body is invalid")
		logger.Error("Request body: ", c.Request.Body)
		logger.Error("Expected body: ", "{username: string, email: string, name_tag: string}")

		c.JSON(http.StatusBadRequest,
			models.Error(
				types.Http.BadRequest(),
				"conflict",
				"Email is already in use",
				"Expected body: {username: string, email: string}",
			),
		)
		return
	}

	email := body.Email
	username := body.Username
	nametag := body.NameTag

	if db.User.GetUserByEmail(email).IsOk() {
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

	if db.Auth.GetConsumedCodeByEmail(email).IsErr() {
		c.JSON(http.StatusUnauthorized,
			models.Error(
				types.Http.Unauthorized(),
				"unauthorized",
				"could not authenticate user",
				"invalid token",
			),
		)
		return
	}

	usernameUsable := db.User.AvailableUsername(username)
	nametagUsable := db.User.AvailableNameTag(nametag)

	if !usernameUsable || !nametagUsable {
		c.JSON(http.StatusBadRequest,
			models.Error(
				types.Http.BadRequest(),
				"bad_request",
				"Username or nametag is already in use",
				strconv.Itoa(utils.BoolToFlags(usernameUsable, nametagUsable)),
			),
		)
		return
	}

	newUser := models.UserCreate{
		Email:    email,
		Username: username,
		NameTag:  nametag,
	}

	if result := db.User.CreateUser(newUser); result.IsErr() {
		result := result.Error()
		c.JSON(result.Code.Get(), result)
	} else {
		c.JSON(http.StatusNoContent, result.Value())
	}
}
