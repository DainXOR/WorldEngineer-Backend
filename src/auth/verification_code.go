package auth

import (
	"dainxor/we/configs"
	"dainxor/we/logger"
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
	configs.DB.Where("email = ?", email).First(&codeDB, "created_at >= NOW() - INTERVAL '5 minutes'")

	if codeDB.ID == 0 {
		logger.Error("Failed to verify email: No valid code found for email provided")
		DeleteExpiredCodes(email)

		return false
	}

	res := codeDB.Code == code

	if res {
		logger.Info("Email verified")
		configs.DB.Delete(&codeDB)
	} else {
		logger.Error("Failed to verify email: Code is invalid")
	}

	return res
}

func DeleteCode(email string, code string) {
	var codeDB models.AuthCodeDB
	configs.DB.Where("email = ?", email).Delete(&codeDB)
}

func DeleteExpiredCodes(email string) {
	var codes []models.AuthCodeDB

	logger.Info("Deleting expired code for email: ", email)
	configs.DB.Where("email = ?", email).Find(&codes, "created_at < NOW() - INTERVAL '5 minutes'")

	for _, c := range codes {
		configs.DB.Delete(c)
	}
}
