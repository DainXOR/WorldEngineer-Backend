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

	registry := models.AuthCodeDB{
		Email:     email,
		Code:      sha256Code,
		CreatedAt: configs.DB.Get().NowFunc(),
	}

	configs.DB.Get().Create(&registry)

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
		return types.ResultErr[models.AuthCodeDB](models.ErrorResponse{
			Type:    "bad_request",
			Message: "Email has no valid codes",
		})
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
func (authType) DeleteConsumedCodesByEmail(email string) types.Result[[]models.AuthCodeDB, models.ErrorResponse] {
	var codesDB []models.AuthCodeDB
	configs.DB.Get().Where("email = ? AND consumed_at IS NOT NULL", email).Delete(&codesDB)

	if len(codesDB) == 0 {
		return types.ResultErr[[]models.AuthCodeDB](models.Error(
			types.Http.NotFound(),
			"not_found",
			"Code registry not found",
		))
	}

	return types.ResultOk[[]models.AuthCodeDB, models.ErrorResponse](codesDB)
}

func (authType) ConsumeCodeById(id uint, code string) types.Result[models.AuthCodeDB, models.ErrorResponse] {
	codeDB := Auth.GetCodeById(id)

	if codeDB.IsErr() {
		logger.Error("Failed to consume code: ", codeDB.Error().Message)
		return codeDB
	}

	codeValue := codeDB.Value()

	if codeValue.ConsumedAt.IsZero() || codeValue.CreatedAt.Before(configs.DB.Get().NowFunc().Add(-5*time.Minute)) {
		return types.ResultErr[models.AuthCodeDB](models.Error(
			types.Http.BadRequest(),
			"bad_request",
			"Code is invalid",
			"Code is expired or consumed",
		))
	}

	match := utils.ComparePassword(codeDB.Value().Code, code)

	if !match {
		return types.ResultErr[models.AuthCodeDB](models.Error(
			types.Http.BadRequest(),
			"bad_request",
			"Code is invalid",
			"Code provided does not match",
		))
	}

	Auth.MarkUsedCodeById(codeDB.Value().ID)

	return codeDB
}

func (authType) ConsumeCodeByEmail(email string, code string) types.Result[models.AuthCodeDB, models.ErrorResponse] {
	codeDB := Auth.GetValidCodeByEmail(email)

	if codeDB.IsErr() {
		return codeDB
	}

	valid := utils.ComparePassword(codeDB.Value().Code, code)

	if !valid {
		return types.ResultErr[models.AuthCodeDB](models.Error(
			types.Http.BadRequest(),
			"bad_request",
			"Code is invalid",
		))
	}

	Auth.MarkUsedCodeById(codeDB.Value().ID)

	return codeDB
}
