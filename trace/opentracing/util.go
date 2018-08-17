package opentracing

import (
	"strconv"
)

func atouint16(s string) uint16 {
	v, _ := strconv.ParseUint(s, 10, 16)
	return uint16(v)
}
