package main

func RunCommand(command string) bool {
	status := true

	switch command {

	case "/":
		Print("> /", 1)

	case "test":
		Print("test command")

	case "q":
		Print("q command")
		status = false

	}

	return status

}
