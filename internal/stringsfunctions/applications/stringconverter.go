package applications

import "strings"

func ToLowerCase(text string) string {
	return strings.ToLower(text)
}

func ToUpperCase(text string) string {
	return strings.ToUpper(text)
}

func Invert(text string) string {
	t := []byte(text)
	newText := ""
	count := len(t) - 1

	for i := count; i > -1; i-- {
		newText = newText + string(t[i])
	}
	return newText
}
