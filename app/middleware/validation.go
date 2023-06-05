package middleware

import "github.com/go-playground/validator/v10"

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "min":
		return "min " + fe.Param() + " characters"
		// case "unique":
		// 	return "This field is unique"
	}

	return "Unknown error"
}
