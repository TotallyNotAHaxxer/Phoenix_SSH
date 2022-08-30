package Phoenix_SSH_Utils_Informative

import (
	"fmt"
	"net"
)

func (Conf *SSH_Configuration) Ping(filler string) {
	var h, p string
	h = Conf.Host
	p = Conf.Port
	c, x := net.Dial("tcp", net.JoinHostPort(h, p))
	if x != nil {
		fmt.Println(Color_Bar("\033[38;5;31m-"), " Could not ping the host, got error -> ", x)
	} else {
		defer c.Close()
		fmt.Println("connection made")
	}
}
