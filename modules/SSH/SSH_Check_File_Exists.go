package Phoenix_SSH_Utils_Informative

import (
	"errors"
	"os"
)

func (Config *SSH_Configuration) Check_F(filename string) bool {
	_, x := os.Stat(filename)
	if x == nil {
		return true
	}
	if errors.Is(x, os.ErrNotExist) {
		return false
	}
	return false
}
