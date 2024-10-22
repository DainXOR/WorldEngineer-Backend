package db

import (
	"dainxor/we/configs"
	"dainxor/we/models"
	"dainxor/we/types"
)

type authType struct{}

var Auth authType

func (authType) GetCodeByEmail(email string) types.Result[models.AuthCodeDB, models.ErrorResponse] {
	var codeDB models.AuthCodeDB

	configs.DB.Where("email = ?", email).First(&codeDB, "created_at >= NOW() - INTERVAL '5 minutes'")

	if codeDB.ID == 0 {
		return types.ResultErr[models.AuthCodeDB](models.ErrorResponse{
			Type:    "bad_request",
			Message: "Email is not registered",
		})
	}

	return types.ResultOk[models.AuthCodeDB, models.ErrorResponse](codeDB)
}
func (authType) GetCodeByValue(code string) types.Result[models.AuthCodeDB, models.ErrorResponse] {
	var codeDB models.AuthCodeDB

	configs.DB.Where("code = ?", code).First(&codeDB)

	if codeDB.ID == 0 {
		return types.ResultErr[models.AuthCodeDB](models.ErrorResponse{
			Type:    "bad_request",
			Message: "Code is not valid",
		})
	}

	return types.ResultOk[models.AuthCodeDB, models.ErrorResponse](codeDB)
}

func (authType) CreateCode(email string, code string) types.Result[models.AuthCodeDB, models.ErrorResponse] {
	registry := models.AuthCodeDB{
		Email:     email,
		Code:      code,
		CreatedAt: configs.DB.NowFunc(),
	}

	configs.DB.Create(&registry)

	if registry.ID == 0 {
		return types.ResultErr[models.AuthCodeDB](models.Error(
			types.Http.InternalServerError(),
			"internal",
			"Failed to create code registry",
		))
	}

	return types.ResultOk[models.AuthCodeDB, models.ErrorResponse](registry)
}

func (authType) DeleteCodeByEmail(email string) types.Result[models.AuthCodeDB, models.ErrorResponse] {
	var codeDB models.AuthCodeDB
	configs.DB.Where("email = ?", email).Delete(&codeDB)

	if codeDB.ID == 0 {
		return types.ResultErr[models.AuthCodeDB](models.Error(
			types.Http.NotFound(),
			"not_found",
			"Code registry not found",
		))
	}

	return types.ResultOk[models.AuthCodeDB, models.ErrorResponse](codeDB)
}
