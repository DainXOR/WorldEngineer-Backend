package utils

import (
	"dainxor/we/configs"
	"dainxor/we/models"
	"math/rand"
	"strconv"
)

func GenerateUserTag(username string) string {
	var nameTag string
	invalid := true

	for invalid {
		randomNumber := strconv.Itoa(rand.Intn(65535))
		nameTag = username + "#" + randomNumber
		invalid = UserTagUsed(nameTag)
	}

	return nameTag
}

func UserTagUsed(nameTag string) bool {
	var user models.UserDB
	configs.DB.Where("name_tag = ?", nameTag).First(&user)
	return user.ID != 0
}

func EmailUsed(email string) bool {
	var user models.UserDB
	configs.DB.Where("email = ?", email).First(&user)
	return user.ID != 0
}
