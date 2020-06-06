package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ContinuumLLC/MyLearningProjects/Submission/parking-lot/src/local/processor"
	//parking-lot/src/local/processor
)

var cs processor.ICommandProcessor

func init() {
	fmt.Println("Initializing Command processor...")
	cs = &processor.CommandProcessorImpl{}
	cs.SetPartnerServiceInstance()
	fmt.Println("Initialized")
}

func main() {
	fmt.Println("Please enter commands to be executed")
	fmt.Println(" file => To execute the commands from file_input.txt")
	fmt.Println(" commands list from sample file => To perform different operations")
	fmt.Println(" exit => To exit the console")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text()
		if strings.EqualFold(command, "exit") {
			break
		} else if strings.EqualFold(command, "file") {
			fmt.Println("Executing commands from file file_input.txt")
			cs.ReadInputFile("file_input.txt")
			break
		} else {
			cs.ProcessCommand(command)
		}
	}

	if scanner.Err() != nil {
		fmt.Println("error occured")
	}

	// argsWithoutProg := os.Args[1:]
	// if len(argsWithoutProg) > 0 {
	// 	command := ""
	// 	for _, s := range argsWithoutProg {
	// 		command = command + " " + s
	// 	}
	// 	cs.ProcessCommand(command)
	// } else {
	// 	//process from file
	// 	cs.ReadInputFile("file_input.txt")
	// }

	//cs.ReadInputFile("file_input.txt")
	// cs.ProcessCommand("create_parking_lot 6")
	// cs.ProcessCommand("status")
	// cs.ProcessCommand("park KA-01-HH-1231 White")
	// cs.ProcessCommand("park KA-01-HH-1232 Black")
	// cs.ProcessCommand("park KA-01-HH-1233 White")
	// cs.ProcessCommand("park KA-01-HH-1234 Black")
	// cs.ProcessCommand("park KA-01-HH-1235 White")
	// cs.ProcessCommand("park KA-01-HH-1236 Black")
	// cs.ProcessCommand("park KA-01-HH-1237 White")
	// cs.ProcessCommand("park KA-01-HH-1238 Black")
	// cs.ProcessCommand("status")
	// cs.ProcessCommand("leave 4")
	// cs.ProcessCommand("status")
	// cs.ProcessCommand("registration_numbers_for_cars_with_colour White")
	// cs.ProcessCommand("slot_numbers_for_cars_with_colour White")
	// cs.ProcessCommand("slot_number_for_registration_number KA-01-HH-3141")
}
