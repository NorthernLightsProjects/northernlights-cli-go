package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("< /")
		text, _ := reader.ReadString('\n')

		if strings.Contains(text, "/") {
			fmt.Println("> /")
		}
	}
}
