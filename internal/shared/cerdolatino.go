package shared

import "strings"

func CerdoLatinoPalabra(texto string) string {
	vocales := "aeiou"
	for _, v := range texto {
		letra := string(v)
		if !strings.Contains(vocales, letra) {
			texto = strings.Replace(texto, letra, "", 1)
			texto = texto + letra + "ay"
			break
		}
	}
	return texto
}

func CerdoLatinoFrace(texto string) string {
	return ""
}
