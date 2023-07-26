package shared

import (
	"math/rand"
	"strings"
	"time"
)

func CrearContraseña(tamaño int) string {
	letras := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "/", "-", "+", "*", "#", "%", "$", "_", "@"}

	var contraseña strings.Builder

	rand.Seed(time.Now().Unix())

	for i := 0; i < tamaño; i++ {
		random := rand.Intn(len(letras))
		letra := letras[random]
		contraseña.WriteString(letra)
	}

	return contraseña.String()

}

//arreglo con todos lo caracteres
