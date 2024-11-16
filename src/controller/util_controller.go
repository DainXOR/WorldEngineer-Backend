package controller

import (
	"dainxor/we/base/logger"
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
	logger.Debug("Checking username: ", username)

	if !db.User.ValidUsername(username) {
		logger.Info("Username is not valid")
		c.JSON(http.StatusBadRequest,
			models.Error(
				types.Http.Conflict(),
				"conflict",
				"Username is not valid",
			),
		)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func (utilType) CheckNameTag(c *gin.Context) {
	nameTag := c.Param("nametag")

	if !db.User.ValidNameTag(nameTag) {
		c.JSON(http.StatusBadRequest,
			models.Error(
				types.Http.BadRequest(),
				"bad_request",
				"Name tag is not valid",
			),
		)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func (utilType) GetProfilePicture(c *gin.Context) {
	id := c.Param("id")
	logger.Debug("Getting profile picture for user: ", id)

	//result := db.User.GetProfilePicture(id)
	//
	//if result.IsErr() {
	//	err := result.Error()
	//	c.JSON(err.Code.AsInt(), err)
	//	return
	//}
	//
	//c.JSON(http.StatusOK,
	//	gin.H{
	//		"picture": result.Value(),
	//	},
	//)
	if id == "9" {
		c.JSON(http.StatusOK, gin.H{"picture": "https://i.pinimg.com/474x/8a/87/72/8a8772480da9d8ae8de32eb20fecb725.jpg"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{})
	}

}
