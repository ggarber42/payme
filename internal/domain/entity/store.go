package entity

import (
	"fmt"
	"regexp"

	"github.com/ggarber42/payme/internal/common/validator"
	"github.com/ggarber42/payme/internal/utils"
)

type StoreData struct {
	Cnpj    string `json:"cnpj"`
	Address *Address
}

var (
	cnpjRegex = regexp.MustCompile(`^\d{2}\.\d{3}\.\d{3}/\d{4}-\d{2}$`)
	cepRegex  = regexp.MustCompile(`^\d{5}-\d{3}$`)
)

type Address struct {
	Cep    string `json:"cep"`
	Street string `json:"street"`
	Number string `json:"number"`
	Other  string `json:"other"`
}

func (sd *StoreData) Validate(v *validator.Validator) {
	v.Check(sd.Address != nil, utils.Errors.MissingField.Key, "missing required field from store")
	v.Check(cnpjRegex.MatchString(sd.Cnpj), utils.Errors.InvalidField.Key, fmt.Sprintf("cpnj %s", utils.Errors.InvalidField.Message))
	v.Check(cepRegex.MatchString(sd.Address.Cep), utils.Errors.InvalidField.Key, fmt.Sprintf("cep %s", utils.Errors.InvalidField.Message))
	v.Check(sd.Address.Street != "", utils.Errors.InvalidField.Key, fmt.Sprintf("street %s", utils.Errors.InvalidField.Message))
	v.Check(sd.Address.Number != "", utils.Errors.InvalidField.Key, fmt.Sprintf("street %s", utils.Errors.InvalidField.Message))
}
