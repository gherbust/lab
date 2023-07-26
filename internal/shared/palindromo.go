package shared

import (
	"strings"

	"github.com/gherbust-meli/lab/internal/stringfuntions/applications"
)

func EsPalindromo(texto string) bool {
	texto = strings.ToLower(texto)
	texto = strings.TrimSpace(texto)
	invertido := applications.Invert(texto)

	return invertido == texto
}
