package Phoenix_SSH_Utils_Informative

import (
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type SSH_Configuration struct {
	Port                                   string     // This will be used to set the current session port
	User                                   string     // This will be used to set the current session username
	Host                                   string     // This will be used to set the current session hostname
	Pass                                   string     // This will be used to set the current session password
	Conn                                   string     // This will be used to set the connection method EG TCP
	Auth                                   string     // This will be used to set authentication methods for SSH TTY `set auth <password/pkey>`
	Pkey                                   ssh.Signer // This will be loaded via file using the load module `load pkey`
	TTY_Height                             int        // This will set the TTY interactive SSH shells defualt height
	TTY_Width                              int        // This will set the TTY interactive SSH shells defualt width
	TTY_Name                               string     // This will set the TTY interactive SSH shells defualt type
	TTY_Mode                               uint32     // This will set the TTY interactive SSH shells ECHO mode number
	SSH_Dst_Directory_For_File_Upload      string
	SSH_Dst_Directory_For_File_Download    string
	SSH_Dst_Directory_For_Directory_Upload string
	*sftp.Client
}

/*
load infofile Config.yaml
set intd /home/reaper/Desktop
utils uloads /home/xea43p3x/Desktop/test-dir/dirdir
*/
