package Phoenix_SSH_Utils_Informative

import (
	"fmt"
	"io/ioutil"
	E "main/modules/Errors"
	Phoenix_Errors "main/modules/Errors"
	Phoenix_Helper "main/modules/Helpers"
	"os"

	"gopkg.in/yaml.v2"
)

type Structure struct {
	Password                  string `yaml:"SSH_Password"`
	Username                  string `yaml:"SSH_Username"`
	Hostname                  string `yaml:"SSH_Hostname"`
	Portname                  string `yaml:"SSH_Portnumb"`
	Connectt                  string `yaml:"SSH_Dial_method"`
	Auth_Type                 string `yaml:"SSH_TTY_Auth_TYPE"`
	TTY_Height                string `yaml:"SSH_TTY_Height"`
	TTY_Width                 string `yaml:"SSH_TTY_Width"`
	TTY_ECHO_MODE             string `yaml:"SSH_TTY_ECHO_Mode"`
	TTY_Name                  string `yaml:"SSH_TTY_Name"`
	SSH_Upload_Directory      string `yaml:"SSH_Dst_Directory_For_File_Upload"`
	Upload_Filepath_name      string
	Upload_File_size          int64
	Upload_File_Permissions   string
	SSH_Download_Directory    string `yaml:"SSH_Dst_Directory_For_File_Download"`
	Download_Filepath_name    string
	Download_File_size        int64
	Download_File_Permissions string
}

func (Config *SSH_Configuration) Reader_SSH_Config(filename string) {

	f, x := ioutil.ReadFile(filename)
	if E.EM(x) {
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Error when trying to read the YAML file -> ", x)
	}
	d := make(map[string]Structure)
	x = yaml.Unmarshal(f, &d)
	if Phoenix_Errors.EM(x) {
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Error when trying to parse the YAML structure -> ", x)
	}
	for _, v := range d {
		Config.Pass = v.Password
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> Config :: \033[38;5;43mNew SSH Password    \033[37m-> ", Config.Pass)
		Config.User = v.Username
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> Config :: \033[38;5;43mNew SSH Username    \033[37m-> ", Config.User)
		Config.Host = v.Hostname
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> Config :: \033[38;5;43mNew SSH Hostname    \033[37m-> ", Config.Host)
		Config.Port = v.Portname
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> Config :: \033[38;5;43mNew SSH Port        \033[37m-> ", Config.Port)
		Config.Conn = v.Connectt
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> Config :: \033[38;5;43mNew SSH Ping Method \033[37m-> ", Config.Conn)
		Config.Auth = v.Auth_Type
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> Config :: \033[38;5;43mNew SSH Auth Method \033[37m-> ", Config.Pass)
		f, _ := Phoenix_Helper.Ret_Int(v.TTY_Height)
		Config.TTY_Height = f
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> Config :: \033[38;5;43mNew SSH TTY Height  \033[37m-> ", Config.TTY_Height)
		f2, _ := Phoenix_Helper.Ret_Int(v.TTY_Width)
		Config.TTY_Width = f2
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> Config :: \033[38;5;43mNew SSH TTY Height  \033[37m-> ", Config.TTY_Width)
		f3 := Phoenix_Helper.Ret_Uint32(v.TTY_ECHO_MODE)
		Config.TTY_Mode = f3
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> Config :: \033[38;5;43mNew SSH TTY ECHO Mode \033[37m-> ", Config.TTY_Name)
		Config.TTY_Name = v.TTY_Name
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> Config :: \033[38;5;43mNew SSH TTY Name    \033[37m-> ", Config.TTY_Name)
		Config.SSH_Dst_Directory_For_File_Upload = v.SSH_Upload_Directory
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> Config :: \033[38;5;43mNew SSH DST File for upload    \033[37m-> ", Config.SSH_Dst_Directory_For_File_Upload)
		if Config.Check_F(v.SSH_Download_Directory) {
			s, x := os.Stat(v.SSH_Download_Directory)
			if x != nil {
				fmt.Println("\033[31m[-] FILE ERROR -> ", x)
			} else {
				v.Download_File_Permissions = s.Mode().Perm().String()
				v.Download_File_size = s.Size()
				v.Download_Filepath_name = s.Name()
				Config.SSH_Dst_Directory_For_File_Download = v.SSH_Download_Directory
				fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> Config :: \033[38;5;43mNew SSH DST File for Download    \033[37m-> ", Config.SSH_Dst_Directory_For_File_Download)
				fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31mFile Info \033[38;5;55m\t===> \033[38;5;43mNew SSH DST File for Download perm    \033[37m-> ", v.Download_File_Permissions)
				fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31mFile Info \033[38;5;55m\t===>  \033[38;5;43mNew SSH DST File for Download size    \033[37m-> ", v.Download_File_size)
				fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31mFile Info \033[38;5;55m\t===> \033[38;5;43mNew SSH DST File for Download name    \033[37m-> ", v.Download_Filepath_name)
			}
		} else {
			fmt.Println("\033[31mCONFIG ERROR: Could not set a defualt upload directory / file for SSH upload, file does not exist")
			fmt.Println("\033[31mCONFIG ERROR: (IGNORE: if you are not using file upload this is not an issue)")
		}

	}
}
