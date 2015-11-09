package main

import "fmt"

func ShowHelp() {
	fmt.Println("usage: ntwrk <command> [options]")
	fmt.Println("commands:")
	commands := [3]string{"help", "info", "ip"}
	descriptions := [3]string{"Show help message for a command",
		"Show detailed network information",
		"\tShow external IPv4 address"}
	for i, cmd := range commands {
		fmt.Printf("    %v\t%v\n", cmd, descriptions[i])
	}
}
