package text

import (
	"strings"
)

func Random(characters string, number int) *string {
	var generated strings.Builder
	for i := 0; i < number; i++ {
		random := Rand.Intn(len(characters))
		randomChar := characters[random]
		generated.WriteString(string(randomChar))
	}

	var str = generated.String()
	return &str
}

var RandomSet = struct {
	Num           string
	MixedAlphaNum string
	UpperAlpha    string
	UpperAlphaNum string
}{
	"0123456789",
	"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	"0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ",
}
