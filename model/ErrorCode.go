package model

type ErrorCode struct {
	Code    int
	Message string
}

func ErrorCodeOf(code int, message string) ErrorCode {
	return ErrorCode{
		Code:    code,
		Message: message,
	}
}
