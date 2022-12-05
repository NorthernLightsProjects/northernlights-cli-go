package main

import (
	"fmt"
	"strings"
)

var (
	// Determines the running state of the loop
	status = true
	// A map of functions.
	commands = map[string]func(params []string){
		"abc": func(params []string) {
			fmt.Println("abc command, params: ", params)
		},
		"/": func(params []string) {
			Print("> /", 2)
		},
		"test": func(params []string) {
			Print("test command", 1)
		},
		"q": func(params []string) {
			Print("q command", 1)
			status = false
		},
		// A function that takes a slice of strings as a parameter and returns a string.
		"echo": func(params []string) {
			text := strings.Join(params, " ")
			Print(text)
		},
		// Clearing the screen.
		"clear": func(params []string) {
			for i := 0; i < 100; i++ {
				Print("")
			}
		},
	}
)

// RunCommand takes a string and a slice of strings as parameters and returns a boolean
func RunCommand(command string, params []string) bool {

	if cmd, ok := commands[command]; ok {
		cmd(params)
	}

	return status
}
