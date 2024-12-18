package db

import (
	"crypto/rand"
	"dainxor/we/base/configs"
	"dainxor/we/base/logger"
	"dainxor/we/models"
	"dainxor/we/types"
	"dainxor/we/utils"
	"math/big"
	"time"
)

type authType struct{}

var Auth authType

func (authType) GenerateCode() types.Result[string, models.ErrorResponse] {
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

func (authType) CreateCode(email string) types.Result[models.AuthCodeDB, models.ErrorResponse] {
	res := Auth.GenerateCode()
	if res.IsErr() {
		return types.ResultErr[models.AuthCodeDB](res.Error())
	}

	resultCode := Auth.SaveCode(email, res.Value())

	return resultCode
}
func (authType) SaveCode(email string, code string) types.Result[models.AuthCodeDB, models.ErrorResponse] {
	sha256Code, err := utils.HashPassword(code)
	if err != nil {
		return types.ResultErr[models.AuthCodeDB](models.Error(
			types.Http.InternalServerError(),
			"internal",
			"Failed to hash code",
		))
	}

	registry := models.AuthCode(sha256Code, email)

	logger.Debug("Code: ", registry)

	configs.DB.Get().Create(&registry)

	logger.Debug("Code ID: ", registry.ID)
	logger.Debug("Created code: ", registry)

	if registry.ID == 0 {
		return types.ResultErr[models.AuthCodeDB](models.Error(
			types.Http.InternalServerError(),
			"internal",
			"Failed to create code registry",
		))
	}

	return types.ResultOk[models.AuthCodeDB, models.ErrorResponse](registry)
}

func (authType) GetCodeById(id uint) types.Result[models.AuthCodeDB, models.ErrorResponse] {
	var codeDB models.AuthCodeDB

	configs.DB.Get().First(&codeDB, id)

	if codeDB.ID == 0 {
		return types.ResultErr[models.AuthCodeDB](models.ErrorResponse{
			Type:    "not_found",
			Message: "Code not found",
		})
	}

	return types.ResultOk[models.AuthCodeDB, models.ErrorResponse](codeDB)
}
func (authType) GetValidCodeByEmail(email string) types.Result[models.AuthCodeDB, models.ErrorResponse] {
	var codeDB models.AuthCodeDB

	configs.DB.Get().Where("email = ?", email).First(&codeDB, "created_at >= NOW() - INTERVAL '5 minutes' AND consumed_at IS NULL")

	if codeDB.ID == 0 {
		return types.ResultErr[models.AuthCodeDB](models.Error(
			types.Http.BadRequest(),
			"bad_request",
			"Email has no valid codes",
		))
	}

	return types.ResultOk[models.AuthCodeDB, models.ErrorResponse](codeDB)
}
func (authType) GetConsumedCodeByEmail(email string) types.Result[models.AuthCodeDB, models.ErrorResponse] {
	var codeDB models.AuthCodeDB

	configs.DB.Get().Where("email = ? AND consumed_at IS NOT NULL", email).First(&codeDB, "consumed_at >= NOW() - INTERVAL '5 minutes'")

	if codeDB.ID == 0 {
		return types.ResultErr[models.AuthCodeDB](models.ErrorResponse{
			Type:    "bad_request",
			Message: "Email has no valid verified codes",
		})
	}

	return types.ResultOk[models.AuthCodeDB, models.ErrorResponse](codeDB)
}
func (authType) GetAllCodesByEmail(email string) types.Result[[]models.AuthCodeDB, models.ErrorResponse] {
	var codesDB []models.AuthCodeDB

	configs.DB.Get().Where("email = ?", email).Find(&codesDB)

	if len(codesDB) == 0 {
		return types.ResultErr[[]models.AuthCodeDB](models.ErrorResponse{
			Type:    "bad_request",
			Message: "Email has no codes",
		})
	}

	return types.ResultOk[[]models.AuthCodeDB, models.ErrorResponse](codesDB)
}
func (authType) GetExpiredCodes() types.Result[[]models.AuthCodeDB, models.ErrorResponse] {
	var codesDB []models.AuthCodeDB

	configs.DB.Get().Where("created_at < NOW() - INTERVAL '5 minutes'").Find(&codesDB)

	if len(codesDB) == 0 {
		return types.ResultErr[[]models.AuthCodeDB](models.ErrorResponse{
			Type:    "bad_request",
			Message: "No expired codes",
		})
	}

	return types.ResultOk[[]models.AuthCodeDB, models.ErrorResponse](codesDB)
}

func (authType) UpdateCodeById(id uint, codeData models.AuthCodeUpdate) types.Result[models.AuthCodeDB, models.ErrorResponse] {
	var codeDB models.AuthCodeDB

	configs.DB.Get().First(&codeDB, id)

	if codeDB.ID == 0 {
		return types.ResultErr[models.AuthCodeDB](models.ErrorResponse{
			Type:    "not_found",
			Message: "Code not found",
		})
	}

	configs.DB.Get().Model(&codeDB).Updates(codeData)

	return types.ResultOf(codeDB, models.Error(
		types.Http.InternalServerError(),
		"internal",
		"Failed to update code",
	), codeDB.ID != 0)
}
func (authType) MarkUsedCodeById(id uint) types.Result[models.AuthCodeDB, models.ErrorResponse] {
	result := Auth.UpdateCodeById(id, models.AuthCodeUpdate{
		ConsumedAt: configs.DB.Get().NowFunc(),
	})

	return result
}

func (authType) DeleteCodeById(id uint) types.Result[models.AuthCodeDB, models.ErrorResponse] {
	var codeDB models.AuthCodeDB

	configs.DB.Get().First(&codeDB, id)

	if codeDB.ID == 0 {
		return types.ResultErr[models.AuthCodeDB](models.ErrorResponse{
			Type:    "not_found",
			Message: "Code not found",
		})
	}

	configs.DB.Get().Delete(&codeDB)

	return types.ResultOk[models.AuthCodeDB, models.ErrorResponse](codeDB)
}
func (authType) DeleteAllCodesByEmail(email string) types.Result[[]models.AuthCodeDB, models.ErrorResponse] {
	var codesDB []models.AuthCodeDB
	configs.DB.Get().Where("email = ?", email).Delete(&codesDB)

	if len(codesDB) == 0 {
		return types.ResultErr[[]models.AuthCodeDB](models.Error(
			types.Http.NotFound(),
			"not_found",
			"Code registry not found",
		))
	}

	return types.ResultOk[[]models.AuthCodeDB, models.ErrorResponse](codesDB)
}
func (authType) DeleteExpiredCodesByEmail(email string) types.Result[[]models.AuthCodeDB, models.ErrorResponse] {
	var codesDB []models.AuthCodeDB
	configs.DB.Get().Where("email = ? AND created_at < NOW() - INTERVAL '5 minutes'", email).Delete(&codesDB)

	if len(codesDB) == 0 {
		return types.ResultErr[[]models.AuthCodeDB](models.Error(
			types.Http.NotFound(),
			"not_found",
			"Code registry not found",
		))
	}

	return types.ResultOk[[]models.AuthCodeDB, models.ErrorResponse](codesDB)
}
func (authType) DeleteConsumedCodesByEmail(email string) types.Optional[models.ErrorResponse] {

	configs.DB.
		Get().
		Where("email = ? AND consumed_at IS NOT NULL", email).
		Delete(&[]models.AuthCodeDB{})

	result := Auth.GetConsumedCodeByEmail(email)

	if result.IsOk() {
		return types.OptionalOf(
			models.Error(
				types.Http.InternalServerError(),
				"internal",
				"Failed to delete consumed codes",
			),
		)
	}

	return types.OptionalEmpty[models.ErrorResponse]()
}

func (authType) ConsumeCodeById(id uint, code string) types.Result[models.AuthCodeDB, models.ErrorResponse] {
	codeDB := Auth.GetCodeById(id)

	if codeDB.IsErr() {
		logger.Warning("Failed to consume code: ", codeDB.Error().Message)
		return codeDB
	}

	codeValue := codeDB.Value()
	consumed := !codeValue.ConsumedAt.Time.IsZero()
	expired := codeValue.CreatedAt.Before(configs.DB.Get().NowFunc().Add(-5 * time.Minute))

	logger.Debug("Code: ", codeValue.Code)
	logger.Debug("Consumed: ", consumed)
	logger.Debug("Expired: ", expired)

	if consumed || expired {
		logger.Warning("Failed to consume code: Code is invalid")
		return types.ResultErr[models.AuthCodeDB](models.Error(
			types.Http.BadRequest(),
			"bad_request",
			"Code is invalid",
			"Code is expired or consumed",
		))
	}

	logger.Debug("Code: ", codeValue.Code)
	logger.Debug("Provided: ", code)
	match := utils.ComparePassword(codeValue.Code, code)

	if !match {
		logger.Warning("Failed to consume code: No match")
		return types.ResultErr[models.AuthCodeDB](models.Error(
			types.Http.BadRequest(),
			"bad_request",
			"Code is invalid",
			"Code provided does not match",
		))
	}

	Auth.MarkUsedCodeById(codeValue.ID)

	return codeDB
}

func (authType) ConsumeCodeByEmail(email string, code string) types.Result[models.AuthCodeDB, models.ErrorResponse] {
	codeDB := Auth.GetValidCodeByEmail(email)

	if codeDB.IsErr() {
		logger.Warning("Failed to consume code: ", codeDB.Error().Message)
		return codeDB
	}

	logger.Debug("Code: ", codeDB.Value().Code)
	logger.Debug("Provided: ", code)
	match := utils.ComparePassword(codeDB.Value().Code, code)

	if !match {
		logger.Warning("Failed to consume code: No match")
		return types.ResultErr[models.AuthCodeDB](models.Error(
			types.Http.BadRequest(),
			"bad_request",
			"Code is invalid",
		))
	}

	Auth.MarkUsedCodeById(codeDB.Value().ID)

	return codeDB
}
