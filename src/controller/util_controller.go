package controller

import (
	"dainxor/we/db"
	"dainxor/we/models"
	"dainxor/we/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type utilType struct{}

var Util utilType

func (utilType) CreateUsername(c *gin.Context) {
	result := db.User.GenerateUsername()

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	c.JSON(http.StatusOK,
		gin.H{
			"username": result.Value(),
		},
	)
}

func (utilType) CreateNameTag(c *gin.Context) {
	username := c.Param("username")
	result := db.User.CreateNameTag(username)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	c.JSON(http.StatusOK,
		gin.H{
			"name_tag": result.Value(),
		},
	)
}

func (utilType) CheckUsername(c *gin.Context) {
	username := c.Param("username")

	if !db.User.ValidUsername(username) {
		c.JSON(http.StatusBadRequest,
			models.Error(
				types.Http.Conflict(),
				"conflict",
				"Username is not valid",
			),
		)
		return
	}
}

func (utilType) CheckNameTag(c *gin.Context) {
	nameTag := c.Param("nameTag")

	if !db.User.AvailableNameTag(nameTag) {
		c.JSON(http.StatusBadRequest,
			models.Error(
				types.Http.Conflict(),
				"conflict",
				"Name tag is already in use",
			),
		)
		return
	}

	// TODO: Check if name tag is valid

	c.JSON(http.StatusOK, gin.H{})
}
