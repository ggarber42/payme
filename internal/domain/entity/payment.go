package entity

import (
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

	if pr.StoreData.Validate(v); !v.IsValid() {
		return v.GetError()
	}

	if pr.CardData.Validate(v); !v.IsValid() {
		return v.GetError()
	}

	if pr.Vendor.Validate(v); !v.IsValid() {
		return v.GetError()
	}

	return nil
}

func (pr *PaymentRequest) Validate(v *validator.Validator) {
	v.Check(pr.CardData != nil, missingKey, utils.MakeErrorMsg("cardData", missingMsg))
	v.Check(pr.Purchase != nil, missingKey, utils.MakeErrorMsg("purchase", missingMsg))
	v.Check(pr.StoreData != nil, missingKey, utils.MakeErrorMsg("store", missingMsg))
	v.Check(pr.Vendor != nil, missingKey, utils.MakeErrorMsg("vendor", missingMsg))
}

type Payment struct {
	CardData
	Purchase
	StoreData
	Vendor
}
