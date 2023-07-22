package applications

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

func invert(text string) string {
	t := []byte(text)
	newText := ""
	count := len(t) - 1

	for i := count; i > -1; i-- {
		newText = newText + string(t[i])
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
		result = invert(text)
	}

	if result == "" {
		return "", errors.New("error al procesar texto")
	}
	return result, nil
}
