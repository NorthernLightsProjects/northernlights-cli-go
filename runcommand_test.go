package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRunCommand(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	commandsKeys, commandsLen := getCommandNamesAndLen(commands)

	fmt.Println("commands keys:", commandsKeys)

	command := commandsKeys[rand.Intn(commandsLen)]

	params := []string{"abc", "xyz"}

	RunCommand(command, params)
}
