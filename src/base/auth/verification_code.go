package auth

import (
	"dainxor/we/base/configs"
	"dainxor/we/base/logger"
	"dainxor/we/models"
	"dainxor/we/types"
	"dainxor/we/utils"

	"crypto/rand"
	"math/big"
)

func GenerateCode() types.Result[string, models.ErrorResponse] {
	bigIntCode, err := utils.Retry(
		func() (*big.Int, error) {
			return rand.Int(rand.Reader, big.NewInt(999999))
		},
		3,
		"Failed to generate verification code: ",
		"Could not generate verification code: ",
	)

	if err != nil {
		logger.Error(err.Error())
		return types.ResultErr[string](models.Error(
			types.Http.InternalServerError(),
			"internal",
			"Failed to generate verification code",
		))
	}

	code := utils.FillZeros(int(bigIntCode.Int64()), 6)
	return types.ResultOk[string, models.ErrorResponse](code)
}

func VerifyCode(email string, code string) bool {
	var codeDB models.AuthCodeDB
	logger.Info("Verifying email: ", email)
	configs.DataBase.Where("email = ?", email).First(&codeDB, "created_at >= NOW() - INTERVAL '5 minutes'")

	if codeDB.ID == 0 {
		logger.Error("Failed to verify email: No valid code found for email provided")
		DeleteExpiredCodes(email)

		return false
	}

	// codeHash, err := utils.HashPassword(code)
	// if err != nil {
	// 	logger.Error("Failed to verify email: Failed to hash code")
	// 	return false
	// }

	// logger.Info("Comparing code: ", codeHash, " with code in DB: ", codeDB.Code)
	res := utils.ComparePassword(codeDB.Code, code)

	if res {
		logger.Info("Email verified")
		configs.DataBase.Delete(&codeDB)
	} else {
		logger.Error("Failed to verify email: Code is invalid")
	}

	return res
}

func DeleteCode(email string, code string) {
	var codeDB models.AuthCodeDB
	configs.DataBase.Where("email = ?", email).Delete(&codeDB)
}

func DeleteExpiredCodes(email string) {
	var codes []models.AuthCodeDB

	logger.Info("Deleting expired code for email: ", email)
	configs.DataBase.Where("email = ?", email).Find(&codes, "created_at < NOW() - INTERVAL '5 minutes'")

	for _, c := range codes {
		configs.DataBase.Delete(c)
	}
}
