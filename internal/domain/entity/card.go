package entity

import "github.com/ggarber42/payme/internal/common/validator"

type CardData struct {
	Name       string `json:"name"`
	Number     string `json:"-"`
	Token      string `json:"token"`
	Flag       string `json:"flag"`
	Cvv        string `json:"-"`
	ExpireDate string `json:"expireDate"`
}

func (c *CardData) isValid(v *validator.Validator) bool {
	return false
}
