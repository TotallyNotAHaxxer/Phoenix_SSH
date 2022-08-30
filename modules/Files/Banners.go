package Phoenix_SSH_Utils_Graphics

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Open_File(banner_file, color string, cl bool) {
	switch cl {
	case true:
		fmt.Println("\x1b[H\x1b[2J\x1b[3J")
	}
	f, x := os.Open(banner_file)
	if x != nil {
		log.Fatal(x)
	} else {
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			fmt.Println(color, scanner.Text())
		}
	}
}
