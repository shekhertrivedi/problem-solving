package main

type Station struct {
	ExecutionTime    string            `json:"executionTime"`
	StationBeanLists []StationBeanList `json:"stationBeanList"`
}
type StationBeanList struct {
	Altitude              string  `json:"altitude,omitempty"`
	AvailableBikes        int64   `json:"availableBikes,omitempty"`
	AvailableDocks        int64   `json:"availableDocks,omitempty"`
	City                  string  `json:"city,omitempty"`
	ID                    int64   `json:"id,omitempty"`
	LandMark              string  `json:"landMark,omitempty"`
	LastCommunicationTime string  `json:"lastCommunicationTime,omitempty"`
	Latitude              float64 `json:"latitude,omitempty"`
	Location              string  `json:"location,omitempty"`
	Longitude             float64 `json:"longitude,omitempty"`
	PostalCode            string  `json:"postalCode,omitempty"`
	StAddress1            string  `json:"stAddress1,omitempty"`
	StAddress2            string  `json:"stAddress2,omitempty"`
	StationName           string  `json:"stationName,omitempty"`
	StatusKey             int64   `json:"statusKey,omitempty"`
	StatusValue           string  `json:"statusValue,omitempty"`
	TestStation           bool    `json:"testStation,omitempty"`
	TotalDocks            int64   `json:"totalDocks,omitempty"`
}

type Response struct {
	Dockable bool   `json:"dockable"`
	Message  string `json:"message"`
}

type ErrorResponse struct {
	Level            string `json:"level,omitempty"`
	Message          string `json:"message,omitempty"`
	RequestParameter struct {
		Argument struct {
			StationId    int    `json:"stationId,omitempty"`
			SearchString string `json:"searchString,omitempty"`
		} `json:"requestParameters,omitempty"`
		Path string `json:"Path,omitempty"`
	} `json:"requestParameters,omitempty"`
	TimeStamp string `json:"timestamp,omitempty"`
}
