package Phoenix_SSH_Utils_Informative

import (
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/crypto/ssh"
)

func (Conf *SSH_Configuration) Load_PKEY(filename string) {
	pkd, x := ioutil.ReadFile(filename)
	if x != nil {
		log.Fatal(x)
	} else {
		pk, x := ssh.ParsePrivateKey(pkd)
		if x != nil {
			log.Fatal(x)
		} else {
			Conf.Pkey = pk
			fmt.Println(Color_Bar("\033[38;5;43m+"), " Information | \033[38;5;43mSet Private Key as \033[37m-> ", Conf.Pkey)
		}
	}
}
