package utils

import "fmt"

func MakeErrorMsg(key, message string) string {
	return fmt.Sprintf("%s: %s", key, message)
}
