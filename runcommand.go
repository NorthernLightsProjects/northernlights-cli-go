package main

import (
	"fmt"
)

var (
	status   = true // determınes the runnıng state of the loop
	commands = map[string]func(params []string){
		"abc": func(params []string) {
			fmt.Println("abc command, params: ", params)
		},
		"/": func(params []string) {
			Print("> /", 1)
		},
		"test": func(params []string) {
			fmt.Println("test command")
		},
		"q": func(params []string) {
			fmt.Println("q command")
			status = false
		},
	}
)

func RunCommand(command string, params []string) bool {

	if cmd, ok := commands[command]; ok {
		cmd(params)
	}

	return status
}
