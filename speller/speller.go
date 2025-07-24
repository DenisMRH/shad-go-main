//go:build !solution

package speller

import (
	"strings"
)

var units = []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
var tens = []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

func Spell(n int64) string {

	fullNumber := []string{}

	switch {
	case n > 0:
	case n < 0:
		fullNumber = append(fullNumber, "minus")
		n = -n
	default:
		return "zero"
	}

	scales := []struct {
		val  int64
		name string
	}{
		{1_000_000_000, "billion"},
		{1_000_000, "million"},
		{1_000, "thousand"},
	}

	for _, scale := range scales {
		if n >= scale.val {
			fullNumber = append(fullNumber, underThousand(n/scale.val)...)
			fullNumber = append(fullNumber, scale.name)
			n %= scale.val
		}
	}

	if n > 0 {
		fullNumber = append(fullNumber, underThousand(n)...)
	}

	return strings.Join(fullNumber, " ")
}

func underHundred(n int64) string {
	switch {
	case n%10 != 0 && n > 19:
		return tens[n/10] + "-" + units[n%10]
	case n%10 == 0 && n > 19:
		return tens[n/10]
	default:
		return units[n]

	}
}

func underThousand(n int64) []string {
	switch {
	case n < 100:
		return []string{underHundred(n)}
	case n > 99 && n%100 != 0:
		return []string{units[n/100], "hundred", underHundred(n % 100)}
	case n > 99 && n%100 == 0:
		return []string{units[n/100], "hundred"}
	default:
		return nil
	}
}
