package service

import (
	//"parking-lot/src/local/models"
	"github.com/ContinuumLLC/MyLearningProjects/Submission/parking-lot/src/local/models"
)

type IParkingService interface {
	CreateParkingLot(n int) error
	ParkVehicle(vehicle models.Vehicle) error
	LeaveParking(n int) error
	GetCurrentSTatusOfParkingLot() []models.ParkingSlot
	GetRegNumbersOfVehicleByColor(color string) ([]string, error)
	GetSlotNosOfVehicleByColor(color string) ([]int, error)
	GetSlotNoForRegNo(number string) (int, error)
}
