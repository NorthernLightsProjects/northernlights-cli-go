package main

import "fmt"

func getCommandNamesAndLen(commands map[string]func(params []string)) (cmdsKeys []string, cmdsLen int) {
	commandsLen := len(commands)

	fmt.Println("commandsLen: ", commandsLen)

	commandsKeys := make([]string, 0, commandsLen)

	for k := range commands {
		commandsKeys = append(commandsKeys, k)
	}

	return commandsKeys, commandsLen
}
