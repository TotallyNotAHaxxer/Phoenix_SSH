package Phoenix_SSH_Utils_Informative

import (
	E "main/modules/Errors"
	"net"

	"golang.org/x/crypto/ssh"
)

func (Conf *SSH_Configuration) Return_Conn() *ssh.Client {
	SSH_Config := &ssh.ClientConfig{}
	switch Conf.Auth {
	case "password":
		SSH_Config.User = Conf.User
		SSH_Config.HostKeyCallback = ssh.InsecureIgnoreHostKey()
		SSH_Config.Auth = []ssh.AuthMethod{
			ssh.Password(Conf.Pass),
		}
	case "pkey":
		SSH_Config.User = Conf.User
		SSH_Config.HostKeyCallback = ssh.InsecureIgnoreHostKey()
		Client_Conf.Auth = []ssh.AuthMethod{
			ssh.PublicKeys(Conf.Pkey),
		}
	}
	cl, x := ssh.Dial(
		Conf.Conn,
		net.JoinHostPort(Conf.Host, Conf.Port), SSH_Config)
	if E.EM(x) {
		Conf.Return_Values("Got error when authenticating / using current settings to authenticate or create a new client session to the SSH server ", x)
		return cl
	} else {
		return cl
	}
}
