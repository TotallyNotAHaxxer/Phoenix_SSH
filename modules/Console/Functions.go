package Phoenix_Command

import (
	"fmt"
	Phoenix_Main_Constants "main/modules"
	Phoenix_Helper "main/modules/Helpers"
)

// this is not a requirement, current alternative to the mapping
func Clear(dt string) {
	fmt.Println("\x1b[H\x1b[2J\x1b[3J")
}

func Auxilary_Set_Pass(password string) {
	Phoenix_Main_Constants.Structure.Pass = password
	fmt.Println(Color_Bar("\033[38;5;43m+"), "New SSH Pass -> ",
		Phoenix_Main_Constants.Structure.Pass)
}

func Auxilary_Set_User(username string) {
	Phoenix_Main_Constants.Structure.User = username
	fmt.Println(Color_Bar("\033[38;5;43m+"), "New SSH User -> ",
		Phoenix_Main_Constants.Structure.User)
}

func Auxilary_Set_Host(hostname string) {
	Phoenix_Main_Constants.Structure.Host = hostname
	fmt.Println(Color_Bar("\033[38;5;43m+"), "New SSH Host -> ",
		Phoenix_Main_Constants.Structure.Host)
}

func Auxilary_Set_Port(port string) {
	Phoenix_Main_Constants.Structure.Port = port
	fmt.Println(Color_Bar("\033[38;5;43m+"), "New SSH Port -> ",
		Phoenix_Main_Constants.Structure.Port)
}

func Auxilary_Set_Dial(dialmethod string) {
	Phoenix_Main_Constants.Structure.Conn = dialmethod
	fmt.Println(Color_Bar("\033[38;5;43m+"), "New SSH Dial -> ",
		Phoenix_Main_Constants.Structure.Conn)
}

func Auxilary_Set_TTY_Width(width_num string) {
	Conv, x := Phoenix_Helper.Ret_Int(width_num)
	if x != nil {
		fmt.Println("\033[31m[-] Error : Could not convert the string to an integer type")
		fmt.Println("\033[31m[-] Status: AUXILARY MODULE - FAILED | Could not set width")
	} else {
		Phoenix_Main_Constants.Structure.TTY_Width = Conv
		fmt.Println(Color_Bar("\033[38;5;43m+"), "New TTY Width -> ",
			Phoenix_Main_Constants.Structure.TTY_Width)
	}
}

func Auxilary_Set_TTY_Height(Height_num string) {
	Conv, x := Phoenix_Helper.Ret_Int(Height_num)
	if x != nil {
		fmt.Println("\033[31m[-] Error : Could not convert the string to an integer type")
		fmt.Println("\033[31m[-] Status: AUXILARY MODULE - FAILED | Could not set height")
	} else {
		Phoenix_Main_Constants.Structure.TTY_Height = Conv
		fmt.Println(Color_Bar("\033[38;5;43m+"), "New TTY Height -> ",
			Phoenix_Main_Constants.Structure.TTY_Width)
	}
}

func Auxilary_Set_TTY_Name(name string) {
	Phoenix_Main_Constants.Structure.TTY_Name = name
	fmt.Println(Color_Bar("\033[38;5;43m+"), "New SSH TTY Name -> ",
		Phoenix_Main_Constants.Structure.TTY_Name)
}

func Auxilary_Set_TTY_Mode(mode string) {
	m := Phoenix_Helper.Ret_Uint32(mode)
	Phoenix_Main_Constants.Structure.TTY_Mode = m
	fmt.Println(Color_Bar("\033[38;5;43m+"), "New SSH TTY ECHO Mode -> ",
		Phoenix_Main_Constants.Structure.TTY_Name)
}

func Auxilary_Set_Auth_Method(method string) {
	Phoenix_Main_Constants.Structure.Auth = method
	fmt.Println(Color_Bar("\033[38;5;43m+"), "New SSH TTY Auth Method -> ",
		Phoenix_Main_Constants.Structure.Auth)
}

func Auxilary_Set_output_File(filepath string) {
	if c := Phoenix_Main_Constants.Structure.Check_F(filepath); c == true {
		Phoenix_Main_Constants.Structure.SSH_Dst_Directory_For_File_Download = filepath
		fmt.Println(Color_Bar("\033[38;5;43m+"), "New SSH output filepath -> ",
			Phoenix_Main_Constants.Structure.SSH_Dst_Directory_For_File_Download)
	} else {
		fmt.Println(Color_Bar("\033[31m-"), "ERROR: Could not set directory, this directory or file does not exist")
	}
}

func Auxilary_Set_Input_File_For_Upload(filepath string) {
	if c := Phoenix_Main_Constants.Structure.Check_F(filepath); c == true {
		Phoenix_Main_Constants.Structure.SSH_Dst_Directory_For_File_Upload = filepath
		fmt.Println(Color_Bar("\033[38;5;43m+"), "New SSH destination filepath -> ",
			Phoenix_Main_Constants.Structure.SSH_Dst_Directory_For_File_Upload)
	} else {
		fmt.Println(Color_Bar("\033[31m-"), "ERROR: Could not set directory, this directory or file does not exist")
	}
}

func Auxilary_Set_Destination_Directory_for_upload_of_directorys(dir string) {
	Phoenix_Main_Constants.Structure.SSH_Dst_Directory_For_Directory_Upload = dir
	fmt.Println(Color_Bar("\033[38;5;43m+"), "New SSH destination upload of dir filepath -> ",
		Phoenix_Main_Constants.Structure.SSH_Dst_Directory_For_Directory_Upload)
}

func Color_Bar(filler string) string {
	a := fmt.Sprintf("\033[38;5;55m|%s\033[38;5;55m|> ", filler)
	return a
}

func Settings(filler string) {
	fmt.Print("\n")
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> Config :: \033[38;5;43mSSH Password                    \033[37m-> ", Phoenix_Main_Constants.Structure.Pass)
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> Config :: \033[38;5;43mSSH Username                    \033[37m-> ", Phoenix_Main_Constants.Structure.User)
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> Config :: \033[38;5;43mSSH Hostname                  \033[37m-> ", Phoenix_Main_Constants.Structure.Host)
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> Config :: \033[38;5;43mSSH Port                      \033[37m-> ", Phoenix_Main_Constants.Structure.Port)
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> Config :: \033[38;5;43mSSH Ping Method               \033[37m-> ", Phoenix_Main_Constants.Structure.Conn)
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> Config :: \033[38;5;43mSSH Auth Method               \033[37m-> ", Phoenix_Main_Constants.Structure.Auth)
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> Config :: \033[38;5;43mSSH TTY Height                \033[37m-> ", Phoenix_Main_Constants.Structure.TTY_Height)
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> Config :: \033[38;5;43mSSH TTY Width                 \033[37m-> ", Phoenix_Main_Constants.Structure.TTY_Width)
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> Config :: \033[38;5;43mSSH TTY ECHO Mode             \033[37m-> ", Phoenix_Main_Constants.Structure.TTY_Mode)
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> Config :: \033[38;5;43mSSH TTY Name                  \033[37m-> ", Phoenix_Main_Constants.Structure.TTY_Name)
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> Config :: \033[38;5;43mSSH SSH DST Download File     \033[37m-> ", Phoenix_Main_Constants.Structure.SSH_Dst_Directory_For_File_Download)
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> Config :: \033[38;5;43mSSH SSH DST Upload File       \033[37m-> ", Phoenix_Main_Constants.Structure.SSH_Dst_Directory_For_File_Upload)
	fmt.Print("\n")
}
