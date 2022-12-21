package main

import (
	s "strings"
)

var (
	// Determines the running state of the loop
	status = true
	// A map of functions.
	commands = map[string]func(params []string){
		"cmds": func(params []string) {
			var cmds string = s.Join(params, ", ")
			Print("Available commands: "+cmds, 1)
		},
		"abc": func(params []string) {
			var _params string = s.Join(params, ", ")
			Print("abc command, params: "+_params, 1)
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
			text := s.Join(params, " ")
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
		if command == "cmds" {
			commandsKeys, _ := getCommandNamesAndLen(commands)
			cmd(commandsKeys)
		} else {
			cmd(params)
		}
	}

	return status
}
