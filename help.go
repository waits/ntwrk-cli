package main

import "fmt"

func ShowHelp() {
	fmt.Println("usage: ntwrk <command> [arguments]\n")
	fmt.Println("commands:")
	commands := [4]string{"geo [ip]", "help", "info [ip]", "ip"}
	descriptions := [4]string{
		"Show GeoIP data about an IP",
		"Show this help message",
		"Show basic info about an IP",
		"\tShow external IPv4 address"}
	for i, cmd := range commands {
		fmt.Printf("    %v\t%v\n", cmd, descriptions[i])
	}
}
