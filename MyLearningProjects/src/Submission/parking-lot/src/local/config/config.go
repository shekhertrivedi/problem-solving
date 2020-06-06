package config

import (
	"parking-lot/src/local/service"
)

type Config struct {
	CmdProcessor   service.ICommandProcessor
	ParkingService service.IParkingService
}

func InitConfig() Config {
	cmdProcessor := &service.CommandProcessorImpl{}
	parkingService := &service.ParkingServiceImpl{}
	return Config{cmdProcessor, parkingService}
}
