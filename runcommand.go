package main

import (
	"fmt"
)

var (
	status   = true // determınes the runnıng state of the loop
	commands = map[string]func(){
		"abc": func() {
			fmt.Println("abc command")

		},
		"/": func() {
			Print("> /", 1)
		},
		"test": func() {
			fmt.Println("test command")

		},
		"q": func() {
			fmt.Println("q command")
			status = false
		},
	}
)

func RunCommand(command string) bool {

	if cmd, ok := commands[command]; ok {
		cmd()
	}

	return status
}
