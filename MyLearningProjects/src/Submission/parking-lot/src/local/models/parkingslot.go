package models

type ParkingSlot struct {
	ID            int
	ParkedVehicle Vehicle
	IsFree        bool
}
