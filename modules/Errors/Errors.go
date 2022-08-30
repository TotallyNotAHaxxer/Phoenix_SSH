package Phoenix_Errors

import (
	"fmt"
	"log"
)

// Standard error
func E(x error, y string) {
	if x != nil {
		switch y {
		case "logf":
			log.Fatal(x)
		case "logp":
			log.Println(x)
		case "panc":
			panic(x)
		case "print":
			print(x)
		case "println":
			println(x)
		case "printfmt":
			fmt.Print(x)
		case "printlnfmt":
			fmt.Println(x)
		case "printffmt":
			fmt.Printf("%s", x)
		}
	}
}

// Standard Message and error
func EM(x error) bool {
	if x != nil {
		return true
	}
	return false
}
