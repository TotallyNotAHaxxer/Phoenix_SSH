package Phoenix_Command

import (
	"fmt"
	"strings"
)

func Parser(I string, Map_of_choice map[string]interface{}) {
	I = strings.TrimSpace(I)
	dt := Map_of_choice[I]
	if dt == nil {
		fmt.Print("\n")
		fmt.Println("[-] ERROR: (HASH map interface matching) - command not found, string didnt match")
		fmt.Println("[-] IGNORE: if command executed anyway, ignore this error, if the error was real")
		fmt.Println("[-] IGNORE: most likely it would have returned to console, if not no error here")
		fmt.Print("\n")
	} else {
		dt.(func(string))("")

	}
}

func Parser_With_Args(Module string, Map_of_choice map[string]interface{}, Module_Arg string) interface{} {
	Module = strings.TrimSpace(Module)
	dt := Map_of_choice[Module]
	if dt == nil {
	} else {
		dt.(func(string))(Module_Arg)
	}
	return dt
}
