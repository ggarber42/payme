package entity

import (
	"github.com/ggarber42/payme/internal/common/validator"
	"github.com/ggarber42/payme/internal/utils"
)

type Purchase struct {
	PurchaseId  string      `json:"purchaseid"`
	Date        string      `json:"date"`
	CustomerId  string      `json:"customerId"`
	TotalAmount float64     `json:"totalAmount"`
	Currency    string      `json:"currency"`
	Products    []Product   `json:"products"`
	Installments Installments `json:"installments"`
}

type Product struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type Installments struct {
	Number int     `json:"numberOfInstallments"`
	Amount float64 `json:"installmentAmount"`
	Total  float64 `json:"totalInstallmentAmount"`
}

var (
	CURRENCY_OPTIONS = []string{"USD", "BRL"}
)

func (p *Purchase) Validate(v *validator.Validator) {
	v.Check(p.PurchaseId == "", missingKey, utils.MakeErrorMsg("purchaseid", missingMsg))
	v.Check(p.Date == "", missingKey, utils.MakeErrorMsg("date", missingMsg))
	v.Check(p.CustomerId == "", missingKey, utils.MakeErrorMsg("customerid", missingMsg))
	v.Check(p.Currency == "", missingKey, utils.MakeErrorMsg("currency", missingMsg))
	v.Check(v.In(p.Currency, CURRENCY_OPTIONS), invalidKey, utils.MakeErrorMsg("currency", invalidMsg))
	v.Check(len(p.Products) == 0, missingKey, utils.MakeErrorMsg("products", missingMsg))

	for _, product := range p.Products {
		v.Check(product.Id == "", missingKey, utils.MakeErrorMsg("product id", missingMsg))
		v.Check(product.Name == "", missingKey, utils.MakeErrorMsg("product name", missingMsg))
	}

	v.Check(p.Installments.Number == 0, missingKey, utils.MakeErrorMsg("numberOfInstallments", missingMsg))
	v.Check(p.Installments.Amount == 0, missingKey, utils.MakeErrorMsg("installmentAmount", missingMsg))
	v.Check(p.Installments.Total == 0, missingKey, utils.MakeErrorMsg("totalInstallmentAmount", missingMsg))
}
