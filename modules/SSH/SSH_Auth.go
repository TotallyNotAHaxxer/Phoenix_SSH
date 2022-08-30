package Phoenix_SSH_Utils_Informative

import (
	"bytes"
	"fmt"
	"net"

	E "main/modules/Errors"

	"golang.org/x/crypto/ssh"
)

func (Conf *SSH_Configuration) Authenticate_With_Settings(filler string) {
	SSH_Config := &ssh.ClientConfig{
		User: Conf.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(Conf.Pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	cl, x := ssh.Dial(
		Conf.Conn,
		net.JoinHostPort(
			Conf.Host, Conf.Port), SSH_Config)
	if E.EM(x) {
		Conf.Return_Values("Got error when authenticating / using current settings to authenticate or create a new client session to the SSH server ", x)
	} else {
		fmt.Println(Color_Bar("\033[38;5;43m+"), " Status      | \033[38;5;43mWas able to properly authenticate and connect to the server")
		fmt.Println(Color_Bar("\033[38;5;43m+"), " Status      | \033[38;5;43mPassword   used -> ", Conf.Pass)
		fmt.Println(Color_Bar("\033[38;5;43m+"), " Status      | \033[38;5;43mUsername   used -> ", Conf.User)
		fmt.Println(Color_Bar("\033[38;5;43m+"), " Status      | \033[38;5;43mHostname   used -> ", Conf.Host)
		fmt.Println(Color_Bar("\033[38;5;43m+"), " Status      | \033[38;5;43mPort       used -> ", Conf.Port)
		fmt.Println(Color_Bar("\033[38;5;43m+"), " Status      | \033[38;5;43mSSH_Method used -> ", Conf.Conn)
		fmt.Println(Color_Bar("\033[38;5;43m+"), " Information | \033[38;5;43mStarting session... ", Conf.Port)

		session, x1 := cl.NewSession()
		if E.EM(x1) {
			Conf.Return_Values("Got error when connecting or creating a new session with the given authentication settings ", x)
		} else {
			defer session.Close()
			var buffer bytes.Buffer
			session.Stdout = &buffer
			if X := session.Run("echo 'hello'"); X != nil {
				str := "Could not return or run a command through the given clients session command [ echo 'hello' ] FAIL "
				Conf.Return_Values(str, X)
			} else {
				fmt.Println("")
				fmt.Println("------------------------------------")
				fmt.Println("Output of command -> echo 'hello'")
				fmt.Println(buffer.String())
			}
		}

	}

}
