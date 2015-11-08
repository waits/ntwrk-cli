package main

import "fmt"
import "os"

// import "io/ioutil"

const (
	RED     = "\x1b[31m"
	GREEN   = "\x1b[32m"
	YELLOW  = "\x1b[33m"
	BLUE    = "\x1b[34m"
	MAGENTA = "\x1b[35m"
	CYAN    = "\x1b[36m"
	RESET   = "\x1b[0m"
)

func main() {
	var action string
	if len(os.Args) > 1 {
		action = os.Args[1]
	} else {
		action = "help"
	}
	switch action {
	case "help":
		help()
	case "ip":
		ip()
	default:
		fmt.Printf("Unknown command '%v'.\n", action)
	}
}

func printcl(color string, str string, vars ...interface{}) {
	fmt.Printf(color+str+RESET+"\n", vars...)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}