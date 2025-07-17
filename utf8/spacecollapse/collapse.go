//go:build !solution

package spacecollapse

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func CollapseSpaces(input string) string {
	var b strings.Builder
	b.Grow(len(input))
	inSpace := false

	for len(input) > 0 {
		r, size := utf8.DecodeRuneInString(input)
		input = input[size:]
		if r == utf8.RuneError && size == 1 {
			r = '\uFFFD'
		}

		if unicode.IsSpace(r) {
			if !inSpace {
				b.WriteByte(' ')
				inSpace = true
			}
		} else {
			b.WriteRune(r)
			inSpace = false
		}

	}
	return b.String()
}
