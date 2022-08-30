package Phoenix_Helper

import "strconv"

func Ret_Int(s string) (int, error) {
	l, x := strconv.Atoi(s)
	return l, x
}
