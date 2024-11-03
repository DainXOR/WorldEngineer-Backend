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

func (utilType) CreateNameTag(c *gin.Context) {
	username := c.Param("username")
	result := db.User.CreateNameTag(username)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.Get(), err)
		return
	}

	c.JSON(http.StatusOK,
		gin.H{
			"name_tag": result.Value(),
		},
	)
}

func (utilType) AvailableNameTag(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{})
}
