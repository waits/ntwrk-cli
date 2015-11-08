package main

import "fmt"

func help() {
	fmt.Println("usage: ntwrk <command> [options]")
	fmt.Println("commands:")
	commands := [2]string{"help", "ip"}
	descriptions := [2]string{"Show help message for a command", "\tShow external IP address"}
	for i, cmd := range commands {
		fmt.Printf("    %v\t%v\n", cmd, descriptions[i])
	}
}
