package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Init time is", time.Now())

	rand.Seed(time.Now().UnixNano())
	fmt.Println("Random number is", rand.Intn(10))

	//reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner((os.Stdin))

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
