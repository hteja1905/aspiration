package main

import (
	"fmt"
	"strings"
	"unicode"
)

func CapitalizeEveryThirdAlphanumericChar(s string) string {
	s = strings.ToLower(s)
	r := []rune(s)
	counter := 0
	for idx, v := range r {
		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			counter++
			if counter%3 == 0 && unicode.IsLower(v) {
				r[idx] = unicode.ToUpper(r[idx])
			}
		}
	}

	return string(r)
}

func main() {
	fmt.Println(CapitalizeEveryThirdAlphanumericChar("Aspiration.com"))
}
