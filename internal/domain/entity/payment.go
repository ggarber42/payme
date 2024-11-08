package entity

import (
	"fmt"
	"net/http"

	"github.com/ggarber42/payme/internal/common/validator"
	"github.com/ggarber42/payme/internal/utils"
)

type PaymentRequest struct {
	CardData  *CardData  `json:"cardData"`
	Purchase  *Purchase  `json:"purchase"`
	StoreData *StoreData `json:"store"`
	Vendor    *Vendor    `json:"vendor"`
}

func (pr *PaymentRequest) Bind(r *http.Request) error {

	v := validator.NewValidator()

	if pr.Validate(v); !v.IsValid() {
		return v.GetError()
	}

	if pr.Vendor.Validate(v); !v.IsValid() {
		return v.GetError()
	}

	if pr.StoreData.Validate(v); !v.IsValid() {
		return v.GetError()
	}

	return nil
}

func (pr *PaymentRequest) Validate(v *validator.Validator) {
	v.Check(pr.CardData != nil, utils.Errors.MissingField.Key, fmt.Sprintf("%s: cardData", utils.Errors.MissingField.Message))
	v.Check(pr.Purchase != nil, utils.Errors.MissingField.Key, fmt.Sprintf("%s: purchase", utils.Errors.MissingField.Message))
	v.Check(pr.StoreData != nil, utils.Errors.MissingField.Key, fmt.Sprintf("%s: store", utils.Errors.MissingField.Message))
	v.Check(pr.Vendor != nil, utils.Errors.MissingField.Key, fmt.Sprintf("%s: vendor", utils.Errors.MissingField.Message))
}

type Payment struct {
	CardData
	Purchase
	StoreData
	Vendor
}
