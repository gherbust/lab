package aplication

import (
	"errors"
	"strings"
)

func toLowerCase(text string) string {
	return strings.ToLower(text)

}

func toUpperCase(text string) string {
	return strings.ToUpper(text)
}

func Invert(text string) string {
	t := []byte(text)
	newText := ""
	count := len(t) - 1

	for i := count; i > -1; i-- {
		newText = newText + string(t[1])
	}
	return newText
}

func Factory(action, text string) (string, error) {
	result := ""
	switch {
	case action == "lower":
		result = toLowerCase(text)
	case action == "upper":
		result = toUpperCase(text)
	case action == "invert":
		result = Invert(text)
	}

	if result == "" {
		return "", errors.New("Error al procesar el texto")
	}
	return result, nil
}
