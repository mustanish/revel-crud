package helpers

import (
	"github.com/revel/revel"
)

type Error struct {
	Error interface{} `json:"error"`
}

func (error Error) FormatError(errors map[string]*revel.ValidationError) map[string]string {
	errorObject := make(map[string]string)
	for key, value := range errors {
		errorObject[key] = value.Message
	}
	return errorObject
}
