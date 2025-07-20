//go:build !solution

package varfmt

import (
	"fmt"
	"strconv"
	"strings"
)

func Sprintf(format string, args ...interface{}) string {
	var b strings.Builder
	var countArgs int
	b.Grow(2 * len(format))

	for len(format) > 0 {
		r := format[0]
		format = format[1:]
		var nxtr byte
		if len(format) != 0 {
			nxtr = format[0]
		}

		if r == 123 {
			if nxtr == 125 {
				format = format[1:]
				wrToBilder(&b, args[countArgs])
			} else {
				end := strings.IndexByte(format, 125)
				argStr := format[:end]
				format = format[end+1:]
				argInt, _ := strconv.Atoi(argStr)
				wrToBilder(&b, args[argInt])
			}
			countArgs++
		} else {
			b.WriteByte(r)
		}

	}

	return b.String()
}

func wrToBilder(b *strings.Builder, v any) {
	switch typedvalue := v.(type) {
	case string:
		b.WriteString(typedvalue)
	case int:
		b.WriteString(strconv.Itoa(typedvalue))
	default:
		b.WriteString(fmt.Sprint(v))
	}
}
