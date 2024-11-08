package utils

// EXIT
type ExitMessageStruct struct {
	EXIT_SUCCESS int
	EXIT_ERROR   int
}

var ExitMessage = ExitMessageStruct{
	EXIT_SUCCESS: 0,
	EXIT_ERROR:   1,
}

// ERROR
type ErrorDetail struct {
	Key     string
	Message string
}

type ErrorsStruct struct {
	MissingField ErrorDetail
	InvalidField ErrorDetail
}

var Errors = ErrorsStruct{
	MissingField: ErrorDetail{
		Key:     "missing_field",
		Message: "missing required field:",
	},
	InvalidField: ErrorDetail{
		Key: "invalid_field",
		Message: "has an invalid value",
	},
}
