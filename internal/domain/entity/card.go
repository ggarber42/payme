package entity

import (
	"regexp"

	"github.com/ggarber42/payme/internal/common/validator"
	"github.com/ggarber42/payme/internal/utils"
)

type CardData struct {
	Name       string `json:"name"`
	Number     string `json:"-"`
	Token      string `json:"token"`
	Flag       string `json:"flag"`
	Cvv        string `json:"-"`
	ExpireDate string `json:"expireDate"`
}

var (
	tokenRegex      = regexp.MustCompile(`[A-Za-z0-9+/]{20,}={0,2}`)
	expireDateRegex = regexp.MustCompile(`^\d{2}/\d{2,4}$`)
)

func (cd *CardData) Validate(v *validator.Validator) {

	v.Check(cd.Name != "", missingKey, utils.MakeErrorMsg("name", missingMsg))
	v.Check(cd.Token != "", missingKey, utils.MakeErrorMsg("token", missingMsg))
	v.Check(tokenRegex.MatchString(cd.Token), invalidKey, utils.MakeErrorMsg("token", invalidMsg))
	v.Check(cd.Flag != "", missingKey, utils.MakeErrorMsg("flag", missingMsg))
	v.Check(cd.ExpireDate != "", missingKey, utils.MakeErrorMsg("expireDate", missingMsg))
	v.Check(expireDateRegex.MatchString(cd.ExpireDate), invalidKey, utils.MakeErrorMsg("expireDate", invalidMsg))
}
