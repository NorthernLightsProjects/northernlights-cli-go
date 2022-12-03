package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
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

		input := make([]string, 0)

		//text, _ := reader.ReadString('\n')
		text := scanner.Text()

		//input = append(input, text...)
		input = append(input, text)

		command := input[0]

		runCommandStatus := RunCommand(command)

		if !runCommandStatus {
			break
		}
	}
}
