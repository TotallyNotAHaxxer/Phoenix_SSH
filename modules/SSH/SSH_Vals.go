package Phoenix_SSH_Utils_Informative

import "fmt"

func (Conf *SSH_Configuration) Return_Values(error_Msg string, x error) {
	fmt.Println("\033[38;5;55m|\033[31m-\033[38;5;55m|> ", error_Msg, " -> ", x)
	fmt.Println("\033[38;5;55m|\033[31m-\033[38;5;55m|> \033[31m Configuration SSH [Pass]   : ", Conf.Pass)
	fmt.Println("\033[38;5;55m|\033[31m-\033[38;5;55m|> \033[31m Configuration SSH [User]   : ", Conf.User)
	fmt.Println("\033[38;5;55m|\033[31m-\033[38;5;55m|> \033[31m Configuration SSH [Port]   : ", Conf.Port)
	fmt.Println("\033[38;5;55m|\033[31m-\033[38;5;55m|> \033[31m Configuration SSH [Host]   : ", Conf.Host)

}

func Color_Bar(filler string) string {
	a := fmt.Sprintf("\033[38;5;55m|%s\033[38;5;55m|> ", filler)
	return a
}
