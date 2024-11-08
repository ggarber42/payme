package validator

import (
	"encoding/json"
	"errors"
)

type Validator struct {
	Errors map[string]string
}

func NewValidator() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

func (v *Validator) In(value string, list []string) bool {
	for _, el := range list {
		if value == el {
			return true
		}
	}
	return false
}

func (v *Validator) IsValid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) PrintErrorJSON() (string, error) {
	errorJson, err := json.Marshal(v.Errors)
	if err != nil {
		return "", err
	}
	return string(errorJson), nil
}

func (v *Validator) GetError() error {
	errorMsg, err := v.PrintErrorJSON()
	if err != nil {
		return err
	}
	return errors.New(errorMsg)
}
