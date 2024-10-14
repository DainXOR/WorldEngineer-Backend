package models

type ErrorResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Detail  string `json:"detail" default:""`
}
