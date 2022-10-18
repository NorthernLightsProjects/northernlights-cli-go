package main

func RunCommand(command string) {

	switch command {

	case "/":
		Print("> /", 1)

	case "test":
		Print("test command")
	}
}
