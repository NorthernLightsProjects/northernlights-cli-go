package main

import (
	"fmt"
)

func Print(text string, printType ...int) {
	_type := 0

	if len(printType) > 0 {
		_type = printType[0]
	}

	switch _type {
	case 0:
		fmt.Println("> " + text)
	case 1:
		fmt.Println("> /")
	}

	return
}
