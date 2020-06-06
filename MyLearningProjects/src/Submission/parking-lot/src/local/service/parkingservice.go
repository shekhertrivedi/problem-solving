package service

import (
	"errors"
	"fmt"
	//"parking-lot/src/local/models"
	"strings"

	"github.com/ContinuumLLC/MyLearningProjects/Submission/parking-lot/src/local/models"
)

func GetPartnerServiceInstance() IParkingService {
	return &ParkingServiceImpl{}
}

type ParkingServiceImpl struct {
	GparkingSlots []models.ParkingSlot
}

func (p *ParkingServiceImpl) LeaveParking(n int) error {
	if n > 0 && n < len(p.GparkingSlots) {
		p.GparkingSlots[n-1].IsFree = true
		p.GparkingSlots[n-1].ParkedVehicle = nil
	} else {
		fmt.Println("n is invalid")
		return errors.New("n is invalid")
	}
	return nil
}

// ParkVehicle
func (p *ParkingServiceImpl) ParkVehicle(vehicle models.Vehicle) error {

	ok, availableSlot, err := p.getAvailableParking()
	if err != nil {
		return err
	}
	if ok {
		p.GparkingSlots[availableSlot].IsFree = false
		p.GparkingSlots[availableSlot].ParkedVehicle = vehicle
	} else {
		fmt.Println("Sorry, parking lot is full")
		return errors.New("Sorry, parking lot is full")
	}
	return nil
}

func (p *ParkingServiceImpl) getAvailableParking() (bool, int, error) {

	if p.GparkingSlots == nil || len(p.GparkingSlots) == 0 {
		fmt.Println("Error: Parking slots not created")
		return false, 0, errors.New("Parking slots not created")
	}

	for index, val := range p.GparkingSlots {
		if val.IsFree {
			return true, index, nil
		}
	}

	return false, 0, nil
}

// CreateParkingLot
func (p *ParkingServiceImpl) CreateParkingLot(n int) error {

	if n > 0 {
		parkingSlots := make([]models.ParkingSlot, n, n)
		var slot models.ParkingSlot
		for i := 0; i < n; i++ {
			slot = models.ParkingSlot{ID: i + 1, IsFree: true}
			parkingSlots[i] = slot
		}
		p.GparkingSlots = parkingSlots
	} else {
		fmt.Println("Error: n is invalid")
		return errors.New("n is invalid")
	}
	return nil
}

func (p *ParkingServiceImpl) GetCurrentSTatusOfParkingLot() []models.ParkingSlot {
	return p.GparkingSlots
}

func (p *ParkingServiceImpl) GetRegNumbersOfVehicleByColor(color string) ([]string, error) {

	var regNos []string

	if len(strings.TrimSpace(color)) == 0 {
		fmt.Println("Error: Invalid color")
		return regNos, errors.New("Invalid color")
	}

	for _, val := range p.GparkingSlots {
		if !val.IsFree {
			if strings.EqualFold(val.ParkedVehicle.GetColor(), color) {
				regNos = append(regNos, val.ParkedVehicle.GetNumber())
			}
		}
	}

	return regNos, nil
}

func (p *ParkingServiceImpl) GetSlotNosOfVehicleByColor(color string) ([]int, error) {

	var slots []int
	if len(strings.TrimSpace(color)) == 0 {
		fmt.Println("Error: Invalid color")
		return slots, errors.New("Invalid color")
	}

	for _, val := range p.GparkingSlots {
		if !val.IsFree {
			if strings.EqualFold(val.ParkedVehicle.GetColor(), color) {
				slots = append(slots, val.ID)
			}
		}
	}
	return slots, nil

}

func (p *ParkingServiceImpl) GetSlotNoForRegNo(number string) (int, error) {

	if len(strings.TrimSpace(number)) == 0 {
		fmt.Println("Error: Invalid registration number")
		return 0, errors.New("Invalid registration number")
	}

	for _, val := range p.GparkingSlots {
		if !val.IsFree {
			if strings.EqualFold(val.ParkedVehicle.GetNumber(), number) {
				return val.ID, nil
			}
		}
	}
	fmt.Println("Error: Reg no is not parked into this lot")
	return 0, errors.New("Reg no is not parked into this lot")

}
