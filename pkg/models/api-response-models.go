package models

type (
	// ErrorResponse : response error
	ErrorResponse struct {
		Code    uint8
		Message string
	}
	// Authresponse : authorization response service
	Authresponse struct {
		Message   string
		LogedUser AuthUser
	}
	// RegisterResponse : register auth response
	RegisterResponse struct {
		Message string
	}
)
