package utils

import "fmt"

func MakeErrorMsg(key, message string) string {
	return fmt.Sprintf("%s: %s", key, message)
}

func Contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
