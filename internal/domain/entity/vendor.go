package entity

import (
	"github.com/ggarber42/payme/internal/common/validator"
	"github.com/ggarber42/payme/internal/utils"
)

type Vendor string

var (
	VENDOR_OPTIONS = []string{"stone", "cielo", "rede"}
)

func (vd *Vendor) Validate(v *validator.Validator) {
	v.Check(v.In(string(*vd), VENDOR_OPTIONS), utils.Errors.InvalidField.Key, "vendor values suported are: stone, cielo and rede")
}
