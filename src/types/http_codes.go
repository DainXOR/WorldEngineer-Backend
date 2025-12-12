package types

import (
	"fmt"
	"net/http"
)

type HttpCode int
type httpCode struct{}

var Http httpCode

func (c HttpCode) AsInt() int {
	return int(c)
}
func (c HttpCode) AsString() string {
	return fmt.Sprint(c.AsInt())
}
func (c HttpCode) Name() string {
	return http.StatusText(c.AsInt())
}

// 2xx Success
func (httpCode) Ok() HttpCode {
	return http.StatusOK
}
func (httpCode) Created() HttpCode {
	return http.StatusCreated
}
func (httpCode) Accepted() HttpCode {
	return http.StatusAccepted
}
func (httpCode) NoContent() HttpCode {
	return http.StatusNoContent
}
func (httpCode) ResetContent() HttpCode {
	return http.StatusResetContent
}
func (httpCode) PartialContent() HttpCode {
	return http.StatusPartialContent
}
func (httpCode) MultiStatus() HttpCode {
	return http.StatusMultiStatus
}
func (httpCode) AlreadyReported() HttpCode {
	return http.StatusAlreadyReported
}
func (httpCode) IMUsed() HttpCode {
	return http.StatusIMUsed
}

// 3xx Redirection
func (httpCode) MultipleChoices() HttpCode {
	return http.StatusMultipleChoices
}
func (httpCode) MovedPermanently() HttpCode {
	return http.StatusMovedPermanently
}
func (httpCode) Found() HttpCode {
	return http.StatusFound
}
func (httpCode) SeeOther() HttpCode {
	return http.StatusSeeOther
}
func (httpCode) NotModified() HttpCode {
	return http.StatusNotModified
}
func (httpCode) UseProxy() HttpCode {
	return http.StatusUseProxy
}
func (httpCode) TemporaryRedirect() HttpCode {
	return http.StatusTemporaryRedirect
}
func (httpCode) PermanentRedirect() HttpCode {
	return http.StatusPermanentRedirect
}

// 4xx Client Error
func (httpCode) BadRequest() HttpCode {
	return http.StatusBadRequest
}
func (httpCode) Unauthorized() HttpCode {
	return http.StatusUnauthorized
}
func (httpCode) Forbidden() HttpCode {
	return http.StatusForbidden
}
func (httpCode) NotFound() HttpCode {
	return http.StatusNotFound
}
func (httpCode) MethodNotAllowed() HttpCode {
	return http.StatusMethodNotAllowed
}
func (httpCode) Conflict() HttpCode {
	return http.StatusConflict
}
func (httpCode) Gone() HttpCode {
	return http.StatusGone
}
func (httpCode) LengthRequired() HttpCode {
	return http.StatusLengthRequired
}
func (httpCode) PreconditionFailed() HttpCode {
	return http.StatusPreconditionFailed
}
func (httpCode) RequestEntityTooLarge() HttpCode {
	return http.StatusRequestEntityTooLarge
}
func (httpCode) RequestURITooLong() HttpCode {
	return http.StatusRequestURITooLong
}
func (httpCode) UnsupportedMediaType() HttpCode {
	return http.StatusUnsupportedMediaType
}
func (httpCode) RequestedRangeNotSatisfiable() HttpCode {
	return http.StatusRequestedRangeNotSatisfiable
}
func (httpCode) ExpectationFailed() HttpCode {
	return http.StatusExpectationFailed
}
func (httpCode) Teapot() HttpCode {
	return http.StatusTeapot
}
func (httpCode) UnprocessableEntity() HttpCode {
	return http.StatusUnprocessableEntity
}
func (httpCode) Locked() HttpCode {
	return http.StatusLocked
}
func (httpCode) FailedDependency() HttpCode {
	return http.StatusFailedDependency
}
func (httpCode) UpgradeRequired() HttpCode {
	return http.StatusUpgradeRequired
}
func (httpCode) PreconditionRequired() HttpCode {
	return http.StatusPreconditionRequired
}
func (httpCode) TooManyRequests() HttpCode {
	return http.StatusTooManyRequests
}
func (httpCode) RequestHeaderFieldsTooLarge() HttpCode {
	return http.StatusRequestHeaderFieldsTooLarge
}
func (httpCode) UnavailableForLegalReasons() HttpCode {
	return http.StatusUnavailableForLegalReasons
}

// 5xx Server Error
func (httpCode) InternalServerError() HttpCode {
	return http.StatusInternalServerError
}
func (httpCode) NotImplemented() HttpCode {
	return http.StatusNotImplemented
}
func (httpCode) BadGateway() HttpCode {
	return http.StatusBadGateway
}
func (httpCode) ServiceUnavailable() HttpCode {
	return http.StatusServiceUnavailable
}
func (httpCode) GatewayTimeout() HttpCode {
	return http.StatusGatewayTimeout
}
func (httpCode) HTTPVersionNotSupported() HttpCode {
	return http.StatusHTTPVersionNotSupported
}
func (httpCode) VariantAlsoNegotiates() HttpCode {
	return http.StatusVariantAlsoNegotiates
}
func (httpCode) InsufficientStorage() HttpCode {
	return http.StatusInsufficientStorage
}
func (httpCode) LoopDetected() HttpCode {
	return http.StatusLoopDetected
}
func (httpCode) NotExtended() HttpCode {
	return http.StatusNotExtended
}
func (httpCode) NetworkAuthenticationRequired() HttpCode {
	return http.StatusNetworkAuthenticationRequired
}
