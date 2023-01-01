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
	typeArray := []string{"response", "warning", ""}
	typeArrayLen := len(typeArray)
	printType := typeArray[rand.Intn(typeArrayLen)]

	Print(text, printType)
}
