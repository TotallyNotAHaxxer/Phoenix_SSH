package Phoenix_SSH_Utils_Informative

import (
	"fmt"
	E "main/modules/Errors"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
)

var host, port, user, pass, method, auth_method string
var Client_Conf = &ssh.ClientConfig{}

func (Conf *SSH_Configuration) Request_TTY(filler string) {
	host = Conf.Host
	port = Conf.Port
	user = Conf.User
	pass = Conf.Pass
	method = Conf.Conn
	auth_method = Conf.Auth
	switch auth_method {
	case "password":
		Client_Conf.User = user
		Client_Conf.HostKeyCallback = ssh.InsecureIgnoreHostKey()
		Client_Conf.Auth = []ssh.AuthMethod{
			ssh.Password(pass),
		}
	case "pkey":
		Client_Conf.User = user
		Client_Conf.HostKeyCallback = ssh.InsecureIgnoreHostKey()
		Client_Conf.Auth = []ssh.AuthMethod{
			ssh.PublicKeys(Conf.Pkey),
		}
	}
	client, x := ssh.Dial(Conf.Conn,
		net.JoinHostPort(Conf.Host, Conf.Port), Client_Conf)
	if E.EM(x) {
		Conf.Return_Values("Got error when authenticating / using current settings to authenticate or create a new client session to the SSH server ", x)
	} else {
		fmt.Println(Color_Bar("\033[38;5;43m+"), " Status      | \033[38;5;43mWas able to properly authenticate and connect to the server")
		fmt.Println(Color_Bar("\033[38;5;43m+"), " Status      | \033[38;5;43mPassword   used  \033[37m-> ", Conf.Pass)
		fmt.Println(Color_Bar("\033[38;5;43m+"), " Status      | \033[38;5;43mUsername   used  \033[37m-> ", Conf.User)
		fmt.Println(Color_Bar("\033[38;5;43m+"), " Status      | \033[38;5;43mHostname   used  \033[37m-> ", Conf.Host)
		fmt.Println(Color_Bar("\033[38;5;43m+"), " Status      | \033[38;5;43mPort       used  \033[37m-> ", Conf.Port)
		fmt.Println(Color_Bar("\033[38;5;43m+"), " Status      | \033[38;5;43mSSH_Method used  \033[37m-> ", Conf.Conn)
		fmt.Println(Color_Bar("\033[38;5;43m+"), " Information | \033[38;5;43mStarting session \033[37m-> OK+")
	}
	session, x1 := client.NewSession()
	if E.EM(x1) {
		Conf.Return_Values("Got error when connecting or creating a new session with the given authentication settings ", x)
		defer session.Close()
	} else {
		defer session.Close()
		fmt.Println(Color_Bar("\033[38;5;43m+"), " Information | \033[38;5;43mSession Started  \033[37m-> OK+")
		session.Stderr = os.Stderr
		fmt.Println(Color_Bar("\033[38;5;43m+"), " Information | \033[38;5;43mSTERR   Pipped   \033[37m-> OK+")
		session.Stdin = os.Stdin
		fmt.Println(Color_Bar("\033[38;5;43m+"), " Information | \033[38;5;43mSTDIN   Pipped   \033[37m-> OK+")
		session.Stdout = os.Stdout
		fmt.Println(Color_Bar("\033[38;5;43m+"), " Information | \033[38;5;43mSTDOUT  Pipped   \033[37m-> OK+ ")
		x1 = session.RequestPty(
			Conf.TTY_Name,
			Conf.TTY_Height,
			Conf.TTY_Width,
			ssh.TerminalModes{
				ssh.ECHO: Conf.TTY_Mode,
			})
		if E.EM(x1) {
			Conf.Return_Values("Got error when trying to request the given psuedo-terminal / PTY please re elval this issue ", x)
		} else {
			x1 = session.Shell()
			if E.EM(x1) {
				Conf.Return_Values("Got error when trying or attempting to start the session shell / PTY please re eveal this issue ", x)
			} else {
				fmt.Println("----------------------------------------------------------")
				fmt.Print("\n")
				session.Wait()
			}
		}

	}

}
