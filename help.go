package main

import "fmt"

func help() {
	fmt.Println("usage: ntwrk <command> [arguments]\n")
	fmt.Println("commands:")
	commands := [5]string{"geo [ip]", "help", "info [ip]", "ip", "test"}
	descriptions := [5]string{
		"Show GeoIP data about an IP",
		"Show this help message",
		"Show basic info about an IP",
		"\tShow external IPv4 address",
		"Run bandwidth tests"}
	for i, cmd := range commands {
		fmt.Printf("    %v\t%v\n", cmd, descriptions[i])
	}
}
