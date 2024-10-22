package auth

import (
	"dainxor/we/configs"
	"dainxor/we/models"
)

func HasCode(email string) bool {
	var codeDB models.AuthCodeDB
	configs.DB.Where("email = ?", email).First(&codeDB, "created_at >= NOW() - INTERVAL '5 minutes'")

	return codeDB.ID != 0
}
