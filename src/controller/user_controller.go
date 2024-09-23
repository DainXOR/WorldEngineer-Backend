package controller

import (
	"dainxor/we/configs"
	"dainxor/we/logger"
	"dainxor/we/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserCreate(c *gin.Context) {
	var user models.UserDB
	var body models.UserCreate

	if err := c.ShouldBindJSON(&body); err != nil {
		logger.Error(err.Error())
		logger.Error("Failed to create user: JSON request body is invalid")
		logger.Error("Request body: ", c.Request.Body)
		logger.Error("Expected body: ", "{username: string, email: string}")

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Username = body.Username
	user.Email = body.Email
	user.CreatedAt = configs.DB.NowFunc()
	user.UpdatedAt = configs.DB.NowFunc()
	user.IDStatus = 1
	user.StatusTimeStamp = configs.DB.NowFunc()

	configs.DB.Create(&user)

	c.JSON(http.StatusCreated,
		models.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			IDStatus: user.IDStatus,
		})
}

func UserGetAll(c *gin.Context) {
	var users []models.UserDB
	var response []models.UserResponse

	configs.DB.Find(&users)

	for _, user := range users {
		response = append(response, models.UserResponse{ID: user.ID, Username: user.Username, Email: user.Email, IDStatus: user.IDStatus})
	}

	c.JSON(http.StatusOK, response)
}
func UserGetByID(c *gin.Context) {
	var user models.UserDB

	id := c.Param("id")

	configs.DB.First(&user, id)

	c.JSON(http.StatusOK,
		models.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			IDStatus: user.IDStatus,
		})
}
func UserGetByStatusID(c *gin.Context) {
	var users []models.UserDB
	var response []models.UserResponse

	id := c.Param("id")

	configs.DB.Where("id_status = ?", id).Find(&users)

	for _, user := range users {
		response = append(response,
			models.UserResponse{
				ID:       user.ID,
				Username: user.Username,
				Email:    user.Email,
				IDStatus: user.IDStatus,
			})
	}

	c.JSON(http.StatusOK, users)
}

func UserUpdateByID(c *gin.Context) {
	var user models.UserDB
	var body models.UserCreate

	id := c.Param("id")

	if err := c.ShouldBindJSON(&body); err != nil {
		logger.Error(err.Error())
		logger.Error("Failed to update user: JSON request body is invalid")
		logger.Error("Request body: ", c.Request.Body)
		logger.Error("Expected body: ", "{username: string, email: string}")

		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON request body is invalid"})
		return
	}

	configs.DB.First(&user, id)

	user.Username = body.Username
	user.Email = body.Email
	user.UpdatedAt = configs.DB.NowFunc()

	configs.DB.Save(&user)

	c.JSON(http.StatusOK,
		models.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			IDStatus: user.IDStatus,
		})
}

func UserDeleteByID(c *gin.Context) {
	var user models.UserDB
	id := c.Param("id")

	configs.DB.Delete(&user, id)

	c.JSON(http.StatusNoContent, gin.H{})
}
