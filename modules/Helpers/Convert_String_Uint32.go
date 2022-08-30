package Phoenix_Helper

import "strconv"

func Ret_Uint32(s string) uint32 {
	val, _ := strconv.ParseUint(s, 10, 64)
	v := uint32(val)
	return v
}
