package entity

import (
	"regexp"

	"github.com/ggarber42/payme/internal/common/validator"
	"github.com/ggarber42/payme/internal/utils"
)

type StoreData struct {
	Cnpj    string   `json:"cnpj"`
	Address *Address `json:"address"`
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
	v.Check(sd.Address != nil, missingKey, utils.MakeErrorMsg("address", missingMsg))
	v.Check(cnpjRegex.MatchString(sd.Cnpj), missingKey, utils.MakeErrorMsg("cnpj", invalidMsg))
	v.Check(cepRegex.MatchString(sd.Address.Cep), missingKey, utils.MakeErrorMsg("cep", invalidMsg))
	v.Check(sd.Address.Street != "", missingKey, utils.MakeErrorMsg("street", invalidMsg))
	v.Check(sd.Address.Number != "", missingKey, utils.MakeErrorMsg("number", invalidMsg))
}
