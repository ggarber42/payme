package entity

import "github.com/ggarber42/payme/internal/utils"

var (
	missingKey = utils.Errors.MissingField.Key
	missingMsg = utils.Errors.MissingField.Message

	invalidKey = utils.Errors.InvalidField.Key
	invalidMsg = utils.Errors.InvalidField.Message
)
