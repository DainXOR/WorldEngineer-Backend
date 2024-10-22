package utils

import (
	"dainxor/we/configs"
	"dainxor/we/models"
	"math/rand"
)

func UserTagGenerate(username string) string {
	var nameTag string
	invalid := true

	for invalid {
		randomNumber := FillZeros(rand.Intn(65535), 5)
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
