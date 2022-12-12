package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestPrint(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	textArray := []string{"abc", "xyz"}
	textArrayLen := len(textArray)
	text := textArray[rand.Intn(textArrayLen)]
	typeArray := []int{0, 1, 2}
	typeArrayLen := len(typeArray)
	printType := typeArray[rand.Intn(typeArrayLen)]

	Print(text, printType)
}
