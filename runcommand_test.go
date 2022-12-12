package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRunCommand(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	commandsLen := len(commands)

	fmt.Println("commandsLen: ", commandsLen)

	commandsKeys := make([]string, 0, commandsLen)

	for k := range commands {
		commandsKeys = append(commandsKeys, k)
	}

	fmt.Println("commands keys:", commandsKeys)

	command := commandsKeys[rand.Intn(commandsLen)]

	params := []string{"abc", "xyz"}

	RunCommand(command, params)
}
