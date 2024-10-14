package auth

import (
	"dainxor/we/configs"
	"dainxor/we/logger"
	"dainxor/we/models"
	"dainxor/we/utils"

	"crypto/rand"
	"math/big"
)

func GenerateCode() uint {
	offset := uint(100000)

	bigIntCode := utils.RetryOrPanic(
		func() (*big.Int, error) {
			return rand.Int(rand.Reader, big.NewInt(int64(999999-offset)))
		},
		3,
		"Failed to generate verification code: ",
		"Could not generate verification code: ",
	)

	return uint(bigIntCode.Int64()) + offset
}

func VerifyCode(email string, codeToVerify uint) bool {
	var codeDB models.AuthCodeDB
	logger.Info("Verifying email: ", email)
	configs.DB.Where("email = ?", email).First(&codeDB, "created_at >= NOW() - INTERVAL '5 minutes'")

	if codeDB.ID == 0 {
		logger.Error("Failed to verify email: Missing email in database")
		return false
	}

	res := codeDB.Code == codeToVerify

	if res {
		logger.Info("Email verified")
		configs.DB.Delete(&codeDB)
	} else {
		logger.Error("Failed to verify email: Code is invalid")
	}

	return res
}

func DeleteCode(email string, code uint) {
	var codeDB models.AuthCodeDB
	configs.DB.Where("email = ?", email).Delete(&codeDB)
}
