package Phoenix_Command

import Phoenix_Main_Constants "main/modules"

// nag = No argument needed
// this will trigger maps that do not need agruments
var Console_map = map[string]interface{}{
	"cls":      Clear, // clear terminal
	"clear":    Clear, // clear terminal
	"settings": Settings,
	"help":     Help_Commands,
	"cmd":      Console_Commands,
}

var Auxilary = map[string]interface{}{
	"pass": Auxilary_Set_Pass,
	"user": Auxilary_Set_User,
	"host": Auxilary_Set_Host,
	"port": Auxilary_Set_Port,
	"conn": Auxilary_Set_Dial,
	"auth": Auxilary_Set_Auth_Method,
	"thei": Auxilary_Set_TTY_Height,
	"tnam": Auxilary_Set_TTY_Name,
	"twid": Auxilary_Set_TTY_Width,
	"tmod": Auxilary_Set_TTY_Mode,
	"outf": Auxilary_Set_output_File,
	"intf": Auxilary_Set_Input_File_For_Upload,
	"intd": Auxilary_Set_Destination_Directory_for_upload_of_directorys,
}

var Load = map[string]interface{}{
	"pkeyfile": Phoenix_Main_Constants.Structure.Load_PKEY,
	"infofile": Phoenix_Main_Constants.Structure.Reader_SSH_Config,
}

var Session = map[string]interface{}{
	"auth":  Phoenix_Main_Constants.Structure.Authenticate_With_Settings,
	"ping":  Phoenix_Main_Constants.Structure.Ping,
	"start": Phoenix_Main_Constants.Structure.Request_TTY,
}

var Exec = map[string]interface{}{
	"execute": Phoenix_Main_Constants.Structure.Exec,
}

var Utils = map[string]interface{}{
	"dloadf": Phoenix_Main_Constants.Structure.Download_Files, // download files or directories
	"uloadf": Phoenix_Main_Constants.Structure.Upload_File,    // upload a single file
}
