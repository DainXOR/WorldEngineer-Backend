package models

import "dainxor/we/types"

type ErrorResponse struct {
	Code    types.HttpCode `json:"code" default:"500"`
	Type    string         `json:"type" default:"internal"`
	Message string         `json:"message" default:""`
	Detail  string         `json:"detail" default:""`
}

func Error(code types.HttpCode, type_ string, information ...string) ErrorResponse {
	message := ""
	detail := ""

	if len(information) > 0 {
		message = information[0]
	}
	if len(information) > 1 {
		detail = information[1]
	}

	return ErrorResponse{
		Code:    code,
		Type:    type_,
		Message: message,
		Detail:  detail,
	}
}

func ErrorNotFound(information ...string) ErrorResponse {
	return Error(types.Http.NotFound(), "not_found", information...)
}

func ErrorInternal(information ...string) ErrorResponse {
	return Error(types.Http.InternalServerError(), "internal", information...)
}
