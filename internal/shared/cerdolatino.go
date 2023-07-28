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
	palabras := strings.Split(texto, " ")
	var nuevoTexto strings.Builder
	for _, v := range palabras {
		nuevoTexto.WriteString(CerdoLatinoPalabra(v))
		nuevoTexto.WriteString(" ")
	}

	return nuevoTexto.String()
}
