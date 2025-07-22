//go:build !solution

package speller

import (
	"strings"
)

func Spell(n int64) string {
	var b strings.Builder
	switch {
	case n > 0:
	case n < 0:
		b.WriteString("minus ")
		n = -n
	default:
		b.WriteString("zero")
	}

	switch {
	case n < 1000:
		underThousand(&b, n)
	case n > 999 && n < 1000000:
		underThousand(&b, n/1000)
		if n%1000 != 0 {
			b.WriteString(" thousand ")
			underThousand(&b, n%1000)
		} else {
			b.WriteString(" thousand")
		}

	case n > 999999 && n < 1000000000:

		underThousand(&b, n/1000000)
		if n%1000000 != 0 {
			b.WriteString(" million ")
			underThousand(&b, (n/1000)%1000)
			if n%1000 != 0 {
				b.WriteString(" thousand ")
				underThousand(&b, n%1000)
			} else {
				b.WriteString(" thousand")
			}
		} else {
			b.WriteString(" million")
		}
	case n > 999999999 && n < 1000000000000:
		underThousand(&b, n/1000000000)
		if n%1000000000 != 0 {
			b.WriteString(" billion ")
			underThousand(&b, (n/1000000)%1000)
			if n%1000000 != 0 {
				b.WriteString(" million ")
				underThousand(&b, (n/1000)%1000)
				if n%1000 != 0 {
					b.WriteString(" thousand ")
					underThousand(&b, n%1000)
				} else {
					b.WriteString(" thousand")
				}
			} else {
				b.WriteString(" million")
			}
		} else {
			b.WriteString(" billion")
		}
	}

	return b.String()
}

func underTwenty(b *strings.Builder, n int64) {
	switch n {
	case 1:
		b.WriteString("one")
	case 2:
		b.WriteString("two")
	case 3:
		b.WriteString("three")
	case 4:
		b.WriteString("four")
	case 5:
		b.WriteString("five")
	case 6:
		b.WriteString("six")
	case 7:
		b.WriteString("seven")
	case 8:
		b.WriteString("eight")
	case 9:
		b.WriteString("nine")
	case 10:
		b.WriteString("ten")
	case 11:
		b.WriteString("eleven")
	case 12:
		b.WriteString("twelve")
	case 13:
		b.WriteString("thirteen")
	case 14:
		b.WriteString("fourteen")
	case 15:
		b.WriteString("fifteen")
	case 16:
		b.WriteString("sixteen")
	case 17:
		b.WriteString("seventeen")
	case 18:
		b.WriteString("eighteen")
	case 19:
		b.WriteString("nineteen")
	}
}

func underHundred(b *strings.Builder, n int64) {
	switch {
	case n < 20:
		underTwenty(b, n)
	case n > 19 && n < 30:
		b.WriteString("twenty")
	case n > 29 && n < 40:
		b.WriteString("thirty")
	case n > 39 && n < 50:
		b.WriteString("forty")
	case n > 49 && n < 60:
		b.WriteString("fifty")
	case n > 59 && n < 70:
		b.WriteString("sixty")
	case n > 69 && n < 80:
		b.WriteString("seventy")
	case n > 79 && n < 90:
		b.WriteString("eighty")
	case n > 89 && n < 100:
		b.WriteString("ninety")

	}
	if n%10 != 0 && n > 19 {
		b.WriteString("-")
		underTwenty(b, n%10)
	}
}

func underThousand(b *strings.Builder, n int64) {
	switch {
	case n > 99:
		underTwenty(b, n/100)
		b.WriteString(" hundred")
		if n%100 != 0 {
			b.WriteString(" ")
			underHundred(b, n%100)
		}
	case n == 1000:
		b.WriteString(" thousand")
	default:
		underHundred(b, n)
	}

}
