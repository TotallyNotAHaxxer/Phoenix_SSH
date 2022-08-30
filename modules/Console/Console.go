package Phoenix_Command

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
)

func Handelreturncon(c chan os.Signal) {
	signal.Notify(c, os.Interrupt)
	for s := <-c; ; s = <-c {
		switch s {
		case os.Interrupt:
			fmt.Println("[-] Detected interruption")
			os.Exit(0)
		case os.Kill:
			fmt.Println("GOT OS KILL")
			os.Exit(0)
		}
	}
}

func Console() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(Color_Bar("\033[38;5;43m+"), "Console  :>> \033[37m")
		t, _ := reader.ReadString('\n')
		t = strings.Replace(t, "\n", "", -1)
		go Handelreturncon(make(chan os.Signal, 1))
		if strings.Contains(t, "console") {
			I := t[7:]
			I = strings.TrimSpace(I)
			a := Parser_With_Args(I, Console_map, "")
			if a == nil {
				fmt.Println("\033[31m[-] ERROR: Command not found in map, this could be a bug given the map just returned empty, if you are not returned to the terminal IGNORE this message, if you were returned to the command prompt / input, then the command actually did not exist")
			}
		}
		if strings.Contains(t, "set") {
			I := t[4:8]
			I = strings.TrimSpace(I)
			a := Parser_With_Args(I, Auxilary, string(t[9:]))
			if a == nil {
				fmt.Println("\033[31m[-] ERROR: Command not found in map, this could be a bug given the map just returned empty, if you are not returned to the terminal IGNORE this message, if you were returned to the command prompt / input, then the command actually did not exist")
			}
		}
		if strings.Contains(t, "session") {
			I := t[8:]
			I = strings.TrimSpace(I)
			Parser(I, Session)
		}
		if strings.Contains(t, "load") {
			I := t[5:13]
			I = strings.TrimSpace(I)
			a := Parser_With_Args(I, Load, string(t[14:]))
			if a == nil {
				fmt.Println("\033[31m[-] ERROR: Command not found in map, this could be a bug given the map just returned empty, if you are not returned to the terminal IGNORE this message, if you were returned to the command prompt / input, then the command actually did not exist")
			}
		}
		if strings.Contains(t, "utils") {
			I := t[6:12]
			I = strings.TrimSpace(I)
			FP := string(t[13:])
			FP = strings.TrimSpace(FP)
			a := Parser_With_Args(I, Utils, FP)
			if a == nil {
				fmt.Println("\033[31m[-] ERROR: Command not found in map, this could be a bug given the map just returned empty, if you are not returned to the terminal IGNORE this message, if you were returned to the command prompt / input, then the command actually did not exist")
			}
		}
		if strings.Contains(t, "exec") {
			I := t[5:12]
			I = strings.TrimSpace(I)
			FP := string(t[13:])
			FP = strings.TrimSpace(FP)
			a := Parser_With_Args(I, Exec, FP)
			if a == nil {
				fmt.Println("\033[31m[-] ERROR: Command not found in map, this could be a bug given the map just returned empty, if you are not returned to the terminal IGNORE this message, if you were returned to the command prompt / input, then the command actually did not exist")
			}
		}
	}
}
