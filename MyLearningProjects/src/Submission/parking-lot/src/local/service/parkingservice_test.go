package service

import (
	//"parking-lot/src/local/models"
	"strings"
	"testing"

	"github.com/ContinuumLLC/MyLearningProjects/Submission/parking-lot/src/local/models"
)

var serviceImplTest ParkingServiceImpl

func TestCreateParkingPositive(t *testing.T) {
	err := serviceImplTest.CreateParkingLot(4)
	if err != nil {
		t.Errorf("Error occured")
	}
}

func TestCreateParkingNegative(t *testing.T) {
	err := serviceImplTest.CreateParkingLot(-4)
	if err == nil {
		t.Errorf("Error occured")
	}
}

func TestGetCurrentSTatusOfParkingLot(t *testing.T) {
	err := serviceImplTest.CreateParkingLot(4)

	parkingSlots := serviceImplTest.GetCurrentSTatusOfParkingLot()

	if err != nil || len(parkingSlots) != 4 {
		t.Errorf("Error occured")
	}
}

func TestParkVehiclePositive1(t *testing.T) {
	err := serviceImplTest.CreateParkingLot(4)
	testVehicle := models.Car{Color: "White", Number: "1234"}
	err = serviceImplTest.ParkVehicle(testVehicle)

	if err != nil || serviceImplTest.GparkingSlots[0].IsFree {
		t.Errorf("Error occured")
	}
}

func TestParkVehiclePositive2(t *testing.T) {
	err := serviceImplTest.CreateParkingLot(4)
	testVehicle := models.Car{Color: "White", Number: "1234"}
	err = serviceImplTest.ParkVehicle(testVehicle)

	if err != nil || serviceImplTest.GparkingSlots[0].ID != 1 {
		t.Errorf("Error occured")
	}
}

func TestParkVehiclePositive3(t *testing.T) {
	err := serviceImplTest.CreateParkingLot(4)
	testVehicle := models.Car{Color: "White", Number: "1234"}
	err = serviceImplTest.ParkVehicle(testVehicle)

	if err != nil || !strings.EqualFold(serviceImplTest.GparkingSlots[0].ParkedVehicle.GetColor(), "White") {
		t.Errorf("Error occured")
	}
}

func TestParkVehiclePositive4(t *testing.T) {
	err := serviceImplTest.CreateParkingLot(4)
	testVehicle := models.Car{Color: "White", Number: "1234"}
	err = serviceImplTest.ParkVehicle(testVehicle)

	if err != nil || !strings.EqualFold(serviceImplTest.GparkingSlots[0].ParkedVehicle.GetNumber(), "1234") {
		t.Errorf("Error occured")
	}
}

//GetRegNumbersOfVehicleByColor(color string) ([]string, error)
func TestGetRegNumbersOfVehicleByColor1(t *testing.T) {
	err := serviceImplTest.CreateParkingLot(4)
	testVehicle := models.Car{Color: "White", Number: "1234"}
	err = serviceImplTest.ParkVehicle(testVehicle)

	arr, err := serviceImplTest.GetRegNumbersOfVehicleByColor("White")

	if err != nil || len(arr) != 1 {
		t.Errorf("Error occured")
	}
}

func TestGetRegNumbersOfVehicleByColor2(t *testing.T) {
	err := serviceImplTest.CreateParkingLot(4)
	testVehicle := models.Car{Color: "White", Number: "1234"}
	err = serviceImplTest.ParkVehicle(testVehicle)

	arr, err := serviceImplTest.GetRegNumbersOfVehicleByColor("White")

	if err != nil || arr[0] != "1234" {
		t.Errorf("Error occured")
	}
}

func TestGetRegNumbersOfVehicleByColor3(t *testing.T) {
	err := serviceImplTest.CreateParkingLot(4)
	testVehicle := models.Car{Color: "White", Number: "1234"}
	err = serviceImplTest.ParkVehicle(testVehicle)

	_, err = serviceImplTest.GetRegNumbersOfVehicleByColor("")

	if err == nil {
		t.Errorf("Error occured")
	}
}

//GetSlotNosOfVehicleByColor(color string) ([]int, error)
func TestGetSlotNosOfVehicleByColor1(t *testing.T) {
	err := serviceImplTest.CreateParkingLot(4)
	testVehicle := models.Car{Color: "White", Number: "1234"}
	err = serviceImplTest.ParkVehicle(testVehicle)

	arr, err := serviceImplTest.GetSlotNosOfVehicleByColor("White")

	if err != nil || arr[0] != 1 {
		t.Errorf("Error occured")
	}
}

func TestGetSlotNosOfVehicleByColor2(t *testing.T) {
	err := serviceImplTest.CreateParkingLot(4)
	testVehicle := models.Car{Color: "White", Number: "1234"}
	err = serviceImplTest.ParkVehicle(testVehicle)

	arr, err := serviceImplTest.GetSlotNosOfVehicleByColor("White")

	if err != nil || len(arr) != 1 {
		t.Errorf("Error occured")
	}
}

func TestGetSlotNosOfVehicleByColor3(t *testing.T) {
	err := serviceImplTest.CreateParkingLot(4)
	testVehicle := models.Car{Color: "White", Number: "1234"}
	err = serviceImplTest.ParkVehicle(testVehicle)

	_, err = serviceImplTest.GetSlotNosOfVehicleByColor("")

	if err == nil {
		t.Errorf("Error occured")
	}
}

//GetSlotNoForRegNo(number string) (int, error)
func TestGetSlotNoForRegNo1(t *testing.T) {
	err := serviceImplTest.CreateParkingLot(4)
	testVehicle := models.Car{Color: "White", Number: "1234"}
	err = serviceImplTest.ParkVehicle(testVehicle)

	_, err = serviceImplTest.GetSlotNoForRegNo("")

	if err == nil {
		t.Errorf("Error occured")
	}
}

func TestGetSlotNoForRegNo2(t *testing.T) {
	err := serviceImplTest.CreateParkingLot(4)
	testVehicle := models.Car{Color: "White", Number: "1234"}
	err = serviceImplTest.ParkVehicle(testVehicle)

	_, err = serviceImplTest.GetSlotNoForRegNo("1234")

	if err != nil {
		t.Errorf("Error occured")
	}
}

func TestGetSlotNoForRegNo3(t *testing.T) {
	err := serviceImplTest.CreateParkingLot(4)
	testVehicle := models.Car{Color: "White", Number: "1234"}
	err = serviceImplTest.ParkVehicle(testVehicle)

	slotNo, err := serviceImplTest.GetSlotNoForRegNo("1234")

	if err != nil || slotNo != 1 {
		t.Errorf("Error occured")
	}
}

//LeaveParking(n int) error

func TestLeaveParking1(t *testing.T) {
	err := serviceImplTest.CreateParkingLot(4)
	testVehicle := models.Car{Color: "White", Number: "1234"}
	err = serviceImplTest.ParkVehicle(testVehicle)

	err = serviceImplTest.LeaveParking(-1)

	if err == nil {
		t.Errorf("Error occured")
	}
}

func TestLeaveParking2(t *testing.T) {
	err := serviceImplTest.CreateParkingLot(4)
	testVehicle := models.Car{Color: "White", Number: "1234"}
	err = serviceImplTest.ParkVehicle(testVehicle)

	err = serviceImplTest.LeaveParking(1)

	if err != nil || !serviceImplTest.GparkingSlots[0].IsFree {
		t.Errorf("Error occured")
	}
}

func TestLeaveParking3(t *testing.T) {
	err := serviceImplTest.CreateParkingLot(4)
	testVehicle := models.Car{Color: "White", Number: "1234"}
	err = serviceImplTest.ParkVehicle(testVehicle)

	err = serviceImplTest.LeaveParking(1)

	if err != nil || serviceImplTest.GparkingSlots[0].ParkedVehicle != nil {
		t.Errorf("Error occured")
	}
}
