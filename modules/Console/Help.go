package Phoenix_Command

import "fmt"

func Help_Commands(filler string) {
	fmt.Println("\033[31m====== Module [ set ] ")
	fmt.Println("")
	fmt.Println("Description: Set is a module that you will use to set session variables\n",
		"Instead of having to reload the script everytime you want to loginto another server\n",
		"You can rather just use the module set in the console to set the iformation, here is its syntax\n",
		"-------------------------------------------------------------------------------\n",
		"\033[32m- set <pass/user/host/port/conn/auth/thei/tnam/twid/tmod> <value>\n",
		"\033[38;5;55m: set pass <value>\033[37m-> Will set the current password to use for executing and authenticating servers\n",
		"\033[38;5;55m: set user <value>\033[37m-> Will set the current username to use for executing and authenticating servers\n",
		"\033[38;5;55m: set host <value>\033[37m-> Will set the current hostname to use for executing and authenticating servers\n",
		"\033[38;5;55m: set port <value>\033[37m-> Will set the current port to use for connecting, pinging, and authenticating servers\n",
		"\033[38;5;55m: set conn <value>\033[37m-> Will set the connection method Example is `set tcp` which will connect using tcp\n",
		"\033[38;5;55m: set auth <value>\033[37m-> Will set the authentication type to use when starting a new ssh session\n",
		"\033[38;5;55m: set thei <value>\033[37m-> Will set the current shell hight when starting a new ssh session \n",
		"\033[38;5;55m: set tnam <value>\033[37m-> Will set the current shell type when starting a new ssh session EG(linux, xterm)\n",
		"\033[38;5;55m: set twid <value>\033[37m-> Will set the current shell width when starting a new ssh session \n",
		"\033[38;5;55m: set tmod <value>\033[37m-> Will set the current shell ECHO number mode when starting a new ssh session\n",
		"\033[38;5;55m: set outf <value>\033[37m-> Will set the output file for downloading files from a SSH server\n",
		"\033[38;5;55m: set intd <value>\033[37m-> Will set the input file / remote ssh destination filepath for file uploads\n",
		"",
	)
	fmt.Println("\033[31m====== Module [ session ] ")
	fmt.Println()
	fmt.Println("Description: session is a module that you will use to auth, ping, or start a new\n",
		"SSH sesion or even log into remote SSH servers, this module requires no argument but rather\n",
		"preset values which can be initalized using the set module. \n",
		"-------------------------------------------------------------------------------\n",
		"\033[32m- session <auth/ping/start>\n",
		"\033[38;5;55m: session auth  \033[37m-> Will attempt to authenticate the current host with the passwords and usernames\n",
		"\033[38;5;55m: session ping  \033[37m-> Will attempt to ping a ssh server to see if it is alive, and output running a command\n",
		"\033[38;5;55m: session start \033[37m-> Will log into a SSH server and start a interactive command line based shell\n",
		"",
	)
	fmt.Println("\033[31m====== Module [ console ] ")
	fmt.Println()
	fmt.Println("Description: Console based commands, usually used to clear output, load settings etc\n",
		"-------------------------------------------------------------------------------\n",
		"\033[32m- console <cls/clear/settings>\n",
		"\033[38;5;55m: console cls      \033[37m-> Will clear all art, banners, and output from the screen\n",
		"\033[38;5;55m: console clear    \033[37m-> Same thing as console cls\n",
		"\033[38;5;55m: console settings \033[37m-> Will load all session variables (port, host, pass, file, key etc)\n",
		"",
	)
	fmt.Println("\033[31m====== Module [ load ]")
	fmt.Println()
	fmt.Println("Description: load allows you to load customized configuration files and private\n",
		"SSH key files for further commands\n",
		"-------------------------------------------------------------------------------\n",
		"\033[32m- load  <pkeyfile/infofile> <file location or path>\n",
		"\033[38;5;55m: load infofile  \033[37m-> Will load the Config.yaml path with custom settings and use that data for commands\n",
		"\033[38;5;55m: load pkeyfile  \033[37m-> Will load and set the private ssh key file for sessions\n",
		"",
	)
	fmt.Println("\033[31m====== Module [ utils ] ")
	fmt.Println()
	fmt.Println("Description: sub utilities and tools that are designed to make the ssh experience better\n",
		"-------------------------------------------------------------------------------\n",
		"\033[32m- utils <downloadf> <value>\n",
		"\033[38;5;55m: utils dloadf <filepath> \033[37m -> Will download all files to a given / set output directory from a given server directory eg `ssh/pi/home/desktop` and output all files found\n\n",
		"\033[38;5;55m: utils uloadf <filepath> \033[37m -> Will attempt to upload a single file to the remote SSH serverfrom a given source directory eg `home/desktop/file.txt`\n",
		"\033[38;5;55m: utils uloadf <filepath> \033[37m -> NOTE: for `utils uloadf` the destination filepath `set intd` must also be a directory + filename /ssh/pi/whatever/output.txt`\n",
		"",
	)

}

func Console_Commands(f string) {
	fmt.Println("",
		"-------------------------------------------------------------------------------\n",
		"\033[38;5;55mset pass         \033[37m: Sets a session password\n",
		"\033[38;5;55mset user         \033[37m: Sets a session username\n",
		"\033[38;5;55mset host         \033[37m: Sets a session hostname\n",
		"\033[38;5;55mset port         \033[37m: Sets a session port number\n",
		"\033[38;5;55mset conn         \033[37m: Sets a session connection / ping type\n",
		"\033[38;5;55mset outf         \033[37m: Sets a output directory of files that will be downloaded to a file\n",
		"\033[38;5;55mset auth         \033[37m: Sets a session authentication form\n",
		"\033[38;5;55mset thei         \033[37m: Sets a SSH TTY shell height\n",
		"\033[38;5;55mset tnam         \033[37m: Sets a SSH TTY shell type\n",
		"\033[38;5;55mset twid         \033[37m: Sets a SSH TTY shell width\n",
		"\033[38;5;55mset tmod         \033[37m: Sets a SSH TTY shell ECHO MODE number\n",
		"\033[38;5;55mset intd         \033[37m: Sets the SSH remote destination filepath and filename for remote file upload\n",
		"\033[38;5;55msession auth     \033[37m: Attempts to auth a ssh server with current variables\n",
		"\033[38;5;55msession ping     \033[37m: Attempts to connect to a ssh server with current variables\n",
		"\033[38;5;55msession start    \033[37m: Attempts to auth, ping, connect, then start a interactive ssh shell\n",
		"\033[38;5;55mload infofile    \033[37m: Attempts to load a given YAML configuration file and sets settings\n",
		"\033[38;5;55mload pkeyfile    \033[37m: Attempts to load and set a PKEY file for ssh server shell sessions\n",
		"\033[38;5;55mconsole help     \033[37m: Loads all console module details such as commands, and syntax\n",
		"\033[38;5;55mconsole cls      \033[37m: Clears the console of all information and history\n",
		"\033[38;5;55mconsole clear    \033[37m: Clears the console of all information and history\n",
		"\033[38;5;55mconsole cmd      \033[37m: Loads all commands (this menu)\n",
		"\033[38;5;55mutils downloadf  \033[37m: Will download files of a SSH filepaths eg (ssh/Desktop)\n",

		"",
	)
}
