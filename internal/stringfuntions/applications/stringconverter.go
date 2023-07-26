package applications

import (
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
		newText = newText + string(t[i])
	}
	return newText
}

func StringConverter(actions []string, text string) []string {
	results := []string{}

	for _, action := range actions {
		result := ""
		action = strings.ToLower(action)
		switch {
		case action == "lower":
			result = toLowerCase(text)
		case action == "upper":
			result = toUpperCase(text)
		case action == "invert":
			result = Invert(text)
		}
		if result == "" {
			result = "error al procesar texto"
		}
		results = append(results, result)
	}

	return results
}
