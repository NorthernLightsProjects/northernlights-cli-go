package main

import (
	"bufio"
	"fmt"
	"strings"
)

func Prompt(scanner *bufio.Scanner) {
	for {
		fmt.Print("< /")

		scanner.Scan()

		//text, _ := reader.ReadString('\n')
		text := scanner.Text()

		//input = append(input, text...)
		textArray := strings.Split(text, " ")

		command := textArray[0]
		params := textArray[1:]

		runCommandStatus := RunCommand(command, params)

		if !runCommandStatus {
			break
		}
	}
}
