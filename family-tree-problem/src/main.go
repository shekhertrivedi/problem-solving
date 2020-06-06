package main

import (
	"bufio"
	"fmt"
	"os"

	input "github.com/ContinuumLLC/familytreeQues/src/InputCommandProcessor"
)

func init() {
	fmt.Println("Initializing Command processor...")
	input.InitializeRepo()

}

func mainT() {
	input.ReadInputFile("file_input.txt")
}
func main() {
	fmt.Println("Please enter commands to be executed")
	fmt.Println("1. \"file\" to read commands from file_input.txt")
	fmt.Println("2. Enter the commands to be executed")
	fmt.Println("3. \"exit\" to leave")
	fmt.Println("")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text()

		switch command {
		case "exit":
			return
		case "file":
			fmt.Println("Executing commands from file file_input.txt")
			input.ReadInputFile("file_input.txt")
		default:
			cp, err := input.GetCommandProcessor(command)
			if err != nil {
				fmt.Println(fmt.Sprintf("Invalid command. Command %v Error %v", command, err))
			}
			cp.ProcessCommand(command)
		}
	}

	if scanner.Err() != nil {
		fmt.Println("error occured")
	}
}
