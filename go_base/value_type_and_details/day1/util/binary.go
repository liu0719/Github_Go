package util

import (
	"math"
	"strings"
)

func BinaryFormat(n int64) string {
	sb := strings.Builder{}
	c := int64(math.Pow(2, 32))
	for i := 0; i < 32; i++ {
		if n&c != 0 {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
		c >>= 1
	}

	return sb.String()
}
