package Phoenix_SSH_Utils_Informative

import (
	"fmt"
	"log"
	E "main/modules/Errors"
	"net"
	"os"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

var (
	HIGH_PINK   = "\033[38;5;198m"
	HIGH_BLUE   = "\033[38;5;21m"
	now         = time.Now()
	FD          = now.Format("15:04:05")
	total_files = 0
	total_dirs  = 0
)

func (Conf *SSH_Configuration) Download_Files(path string) {

	SDL := len(path)
	LTD := Conf.SSH_Dst_Directory_For_File_Download // Where the downloaded files will appear
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
	c, x := ssh.Dial(
		Conf.Conn,
		net.JoinHostPort(
			Conf.Host, Conf.Port),
		SSH_Config)
	if E.EM(x) {
		Conf.Return_Values("Got error when authenticating / using current settings to authenticate or create a new client session to the SSH server ", x)
	} else {
		fmt.Println(Color_Bar("\033[38;5;43m+"), " Status      | \033[38;5;43mWas able to properly authenticate and connect to the server")
		fmt.Println(Color_Bar("\033[38;5;43m+"), " Status      | \033[38;5;43mPassword   used -> ", Conf.Pass)
		fmt.Println(Color_Bar("\033[38;5;43m+"), " Status      | \033[38;5;43mUsername   used -> ", Conf.User)
		fmt.Println(Color_Bar("\033[38;5;43m+"), " Status      | \033[38;5;43mHostname   used -> ", Conf.Host)
		fmt.Println(Color_Bar("\033[38;5;43m+"), " Status      | \033[38;5;43mPort       used -> ", Conf.Port)
		fmt.Println(Color_Bar("\033[38;5;43m+"), " Status      | \033[38;5;43mSSH_Method used -> ", Conf.Conn)
		defer c.Close()
		fmt.Println(Color_Bar("\033[38;5;43m+"), " Information | \033[31mStarting new SFTP Session...", Conf.Conn)
		cl, x := sftp.NewClient(c)
		if E.EM(x) {
			Conf.Return_Values("Got error when trying to create a new SFTP client using previous output from connection  ", x)
		} else {
			defer cl.Close()
			w := cl.Walk(path)
			for w.Step() {
				if x := w.Err(); x != nil {
					log.Println(x)
				}
				if w.Stat().IsDir() {
					total_dirs++
					fmt.Println(HIGH_BLUE+"["+HIGH_PINK+FD+HIGH_BLUE+"]"+HIGH_BLUE+"["+HIGH_PINK+"Download"+HIGH_BLUE+"] Found directory -> ", w.Stat().Name())
					dir_path := LTD + w.Path()[SDL:]
					if x := os.MkdirAll(dir_path, os.ModePerm); x != nil {
						Conf.Return_Values("Got error when trying to make a new directory given the source directory, during downloading the remote source directory ", x)
					}
				} else {
					total_files++
					fmt.Println(HIGH_BLUE+"["+HIGH_PINK+FD+HIGH_BLUE+"]"+HIGH_BLUE+"["+HIGH_PINK+"Download"+HIGH_BLUE+"] Found Filename  -> ", w.Stat().Name())
					fp := LTD + w.Path()[SDL:]
					f, x := cl.Open(w.Path())
					if x != nil {
						Conf.Return_Values("Got error when trying to open the following path ", x)
					} else {
						defer f.Close()
						of, x := os.OpenFile(
							fp,
							os.O_CREATE|os.O_WRONLY,
							os.ModePerm)
						if x != nil {
							Conf.Return_Values("Got error when trying to open the directory  ", x)
						} else {
							defer of.Close()
						}
						_, x = f.WriteTo(of)
						if x != nil {
							Conf.Return_Values("Got error when writing to the output directory ", x)
						}
					}
				}
			}
			fmt.Println(HIGH_BLUE+"["+HIGH_PINK+FD+HIGH_BLUE+"]"+HIGH_BLUE+"["+HIGH_PINK+"Saver"+HIGH_BLUE+"] \033[37mFound total files       -? ", total_files)
			fmt.Println(HIGH_BLUE+"["+HIGH_PINK+FD+HIGH_BLUE+"]"+HIGH_BLUE+"["+HIGH_PINK+"Saver"+HIGH_BLUE+"] \033[37mFound total direcrtories -? ", total_dirs)
			fmt.Println(HIGH_BLUE+"["+HIGH_PINK+FD+HIGH_BLUE+"]"+HIGH_BLUE+"["+HIGH_PINK+"Saver"+HIGH_BLUE+"] \033[37mFound All data saved to ->  ", Conf.SSH_Dst_Directory_For_File_Download)
		}

	}
}
