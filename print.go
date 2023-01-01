package main

import (
	"fmt"
)

func Print(text string, printType ...string) {
	_type := ""

	if len(printType) > 0 {
		_type = printType[0]
	}

	switch _type {
	case "response":
		fmt.Println("> " + text)
	case "warning":
		fmt.Println("[!]" + text)
	default:
		fmt.Println(text)
	}
}
