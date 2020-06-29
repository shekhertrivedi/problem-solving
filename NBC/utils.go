package main

import (
	"time"
)

func Paginate(pageNum int, pageSize int, sliceLength int) (int, int) {
	start := pageNum * pageSize

	if start > sliceLength {
		start = sliceLength
	}

	end := start + pageSize
	if end > sliceLength {
		end = sliceLength
	}

	return start, end
}

func CheckService(station []StationBeanList, status string) (stationLists []StationBeanList) {

	for _, service := range station {
		if service.StatusValue == status {
			stationLists = append(stationLists, service)
		}
	}
	return stationLists
}

func UpdateErrorMessage(level, msg string, stationid int, search, path string) ErrorResponse {
	errorResp := ErrorResponse{}
	errorResp.Level = level
	errorResp.Message = msg
	errorResp.RequestParameter.Argument.SearchString = search
	errorResp.RequestParameter.Argument.StationId = stationid
	errorResp.RequestParameter.Path = path
	errorResp.TimeStamp = time.Now().Format(time.RFC850)

	return errorResp

}
