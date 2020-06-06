package processor

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	//"parking-lot/src/local/models"
	//"parking-lot/src/local/service"
	"strconv"
	"strings"

	"github.com/ContinuumLLC/MyLearningProjects/Submission/parking-lot/src/local/models"
	"github.com/ContinuumLLC/MyLearningProjects/Submission/parking-lot/src/local/service"
)

type ICommandProcessor interface {
	ProcessCommand(command string) error
	SetPartnerServiceInstance()
	ReadInputFile(fileName string) error
}

type CommandProcessorImpl struct {
	parkingSeriveInstance service.IParkingService
}

func (cp *CommandProcessorImpl) SetPartnerServiceInstance() {
	cp.parkingSeriveInstance = service.GetPartnerServiceInstance()
}

func (cp *CommandProcessorImpl) ProcessCommand(command string) error {
	if len(command) == 0 {
		fmt.Println("Error: Invalid command ")
		return errors.New("Invalid command")
	}
	commands := strings.Fields(command)
	switch commands[0] {
	case "create_parking_lot":
		return CreateCommandProcessor(commands, cp)

	case "park":
		return ParkCommandProcessor(commands, cp)

	case "leave":
		return LeaveCommandProcessor(commands, cp)

	case "status":
		return StatusCommandProcessor(commands, cp)

	case "registration_numbers_for_cars_with_colour":
		return RegistrationNosForCarsByColorCommandProcessor(commands, cp)

	case "slot_numbers_for_cars_with_colour":
		return SlotNosForCarsByColorCommandProcessor(commands, cp)

	case "slot_number_for_registration_number":
		return SlotNosForRegNoCommandProcessor(commands, cp)

	default:
		fmt.Println("Error: Invalid command ")
		return errors.New("Invalid command")
	}
	return nil
}

func CreateCommandProcessor(command []string, cp *CommandProcessorImpl) error {
	i, err := strconv.Atoi(command[1])
	if err != nil {
		return err
	}
	return cp.parkingSeriveInstance.CreateParkingLot(i)
}

func ParkCommandProcessor(command []string, cp *CommandProcessorImpl) error {
	number := command[1]
	color := command[2]
	if len(strings.TrimSpace(number)) == 0 || len(strings.TrimSpace(color)) == 0 {
		fmt.Println("Error: Invalid vehicle details")
		return errors.New("Invalid vehicle details")
	}
	vehicle := models.Car{Color: color, Number: number}

	return cp.parkingSeriveInstance.ParkVehicle(vehicle)
}

func LeaveCommandProcessor(command []string, cp *CommandProcessorImpl) error {
	position, err := strconv.Atoi(command[1])
	if err != nil {
		return err
	}
	if position < 1 {
		fmt.Println("Error: Invalid position")
		return errors.New("Invalid position")
	}

	return cp.parkingSeriveInstance.LeaveParking(position)
}

func StatusCommandProcessor(command []string, cp *CommandProcessorImpl) error {
	fmt.Println(cp.parkingSeriveInstance.GetCurrentSTatusOfParkingLot())
	return nil
}

func RegistrationNosForCarsByColorCommandProcessor(command []string, cp *CommandProcessorImpl) error {
	regNos, err := cp.parkingSeriveInstance.GetRegNumbersOfVehicleByColor(command[1])
	if err != nil {
		return err
	}
	fmt.Println(regNos)
	return nil
}

func SlotNosForCarsByColorCommandProcessor(command []string, cp *CommandProcessorImpl) error {
	slotNos, err := cp.parkingSeriveInstance.GetSlotNosOfVehicleByColor(command[1])
	if err != nil {
		return err
	}
	fmt.Println(slotNos)
	return nil
}

func SlotNosForRegNoCommandProcessor(command []string, cp *CommandProcessorImpl) error {
	slotNos, err := cp.parkingSeriveInstance.GetSlotNoForRegNo(command[1])
	if err != nil {
		return err
	}
	fmt.Println(slotNos)
	return nil

}

func (cp *CommandProcessorImpl) ReadInputFile(fileName string) error {

	if len(fileName) == 0 {
		fmt.Println("Error: Invalid command ")
		return errors.New("Invalid command ")
	}
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cp.ProcessCommand(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
