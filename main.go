package main

import (
	"fmt"
	Phoenix_SSH_Utils_Constants "main/modules"
	Phoenix_SSH_Utils_Color "main/modules/Coloring"
	Phoenix_SSH_Utils_Terminal "main/modules/Console"
	Phoenix_SSH_Utils_Files "main/modules/Files"
)

func main() {
	Phoenix_SSH_Utils_Constants.Structure.TTY_Height = 40
	Phoenix_SSH_Utils_Constants.Structure.TTY_Width = 80
	Phoenix_SSH_Utils_Constants.Structure.TTY_Name = "vt100"
	Phoenix_SSH_Utils_Constants.Structure.TTY_Mode = 0
	Phoenix_SSH_Utils_Constants.Structure.Auth = "password"
	Phoenix_SSH_Utils_Constants.Structure.Conn = "tcp"

	Phoenix_SSH_Utils_Files.Open_File(
		"modules/Text/banner.txt",
		Phoenix_SSH_Utils_Color.GSIVQPMV,
		true)
	fmt.Println(Phoenix_SSH_Utils_Color.XUAXFHWO, "─────────────────────────────────────────────────────────────────────────────────────────")
	Phoenix_SSH_Utils_Terminal.Console()
}
