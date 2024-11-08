package entity_interface

import "github.com/ggarber42/payme/internal/common/validator"

type IEntity interface {
	isValid(v validator.Validator) bool
}
