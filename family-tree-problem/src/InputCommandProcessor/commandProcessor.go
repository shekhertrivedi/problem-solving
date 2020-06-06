package input

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/ContinuumLLC/familytreeQues/src/services"
)

// Processor Processor
type Processor interface {
	ProcessCommand(command string)
}

// InitializeRepo InitializeRepo
func InitializeRepo() {
	services.InitializeRepo()
}

// ReadInputFile ReadInputFile
func ReadInputFile(fileName string) error {

	if len(fileName) == 0 {
		fmt.Println("Error: Invalid commands. Empty file.")
		return errors.New("Invalid command ")
	}
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		command := scanner.Text()
		cp, err := GetCommandProcessor(command)
		if err != nil {
			fmt.Println(fmt.Sprintf("Invalid command. Command %v Error %v", command, err))
		}
		cp.ProcessCommand(command)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
