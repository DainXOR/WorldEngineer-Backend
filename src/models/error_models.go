package models

import (
	"dainxor/we/types"
	"errors"
	"fmt"
)

type ErrorResponse struct {
	Code   types.HttpCode `json:"code"`
	Err    error          `json:"error"`
	Detail string         `json:"detail"`
}

func Error(code types.HttpCode, information ...string) ErrorResponse {
	message := ""
	detail := ""

	if len(information) > 0 {
		message = information[0]
	}
	if len(information) > 1 {
		detail = information[1]
	}

	return ErrorResponse{
		Code:   code,
		Err:    errors.New(message),
		Detail: detail,
	}
}

func (m *ErrorResponse) Error() string {
	return fmt.Sprintf("%s (%s): %s - %s", m.Code.Name(), m.Code.AsString(), m.Err, m.Detail)
}

func ErrorNotFound(information ...string) ErrorResponse {
	return Error(types.Http.NotFound(), information...)
}

func ErrorInternal(information ...string) ErrorResponse {
	return Error(types.Http.InternalServerError(), information...)
}
