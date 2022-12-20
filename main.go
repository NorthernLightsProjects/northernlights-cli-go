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

	Prompt(scanner)
}
