package Phoenix_SSH_Utils_Informative

import (
	"fmt"
	E "main/modules/Errors"
	"os"

	"github.com/pkg/sftp"
)

func (Conf *SSH_Configuration) Upload_File(path_to_send string) {
	client := Conf.Return_Conn()
	defer client.Close()
	sftp, x2 := sftp.NewClient(client)
	if E.EM(x2) {
		fmt.Println("\033[31m[--] ERROR: CLIENT: SFTP: SSH -> Got error when trying to start a new client with the returned configuration and authentication settings -> ", x2)
		defer sftp.Close()
	} else {
		fmt.Println(HIGH_BLUE+"["+HIGH_PINK+FD+HIGH_BLUE+"]"+HIGH_BLUE+"["+HIGH_PINK+"Upload"+HIGH_BLUE+"] Client created ", path_to_send)
		sf, x3 := os.Open(path_to_send)
		if E.EM(x3) {
			fmt.Println("\033[31m[-] ERROR: Could not open the source filepath to send to the remote SSH server, got an error when loading -> ", x3)
			defer sf.Close()
		} else {
			fmt.Println(HIGH_BLUE+"["+HIGH_PINK+FD+HIGH_BLUE+"]"+HIGH_BLUE+"["+HIGH_PINK+"Upload"+HIGH_BLUE+"] Found Filename  -> ", path_to_send)
			defer sf.Close()
			Df, x4 := sftp.Create(Conf.SSH_Dst_Directory_For_File_Upload)
			if E.EM(x4) {
				fmt.Println("\033[31m[-] ERROR: SFTP: Could not create the destination directory for remote SSH file upload -> ", x4)
				defer Df.Close()
			} else {
				fmt.Println(HIGH_BLUE+"["+HIGH_PINK+FD+HIGH_BLUE+"]"+HIGH_BLUE+"["+HIGH_PINK+"Upload"+HIGH_BLUE+"] Created destination directory -> ", Conf.SSH_Dst_Directory_For_File_Upload)
				defer Df.Close()
				if _, err := Df.ReadFrom(sf); err != nil {
					fmt.Println("\033[31m[-] ERROR: FILE: SFTP: Could not read from the source file dirdctory in the destination directory on the server -> ", err)

				} else {
					fmt.Println(HIGH_BLUE+"["+HIGH_PINK+FD+HIGH_BLUE+"]"+HIGH_BLUE+"["+HIGH_PINK+"Status"+HIGH_BLUE+"] File has been uploaded ", path_to_send)
				}
				defer sftp.Close()
			}

		}

	}

}
