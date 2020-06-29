package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Get All Station Data
func GetStations(w http.ResponseWriter, req *http.Request) {
	log.Info("Requested API - ", req.URL.Path)
	var stationResp []StationBeanList
	//Filter only required data
	for _, v := range Cache.StationBeanLists {
		var stationlist StationBeanList
		stationlist.StationName = v.StationName
		stationlist.StAddress1 = v.StAddress1
		stationlist.StAddress2 = v.StAddress2
		stationlist.City = v.City
		stationlist.PostalCode = v.PostalCode
		stationlist.AvailableBikes = v.AvailableBikes
		stationlist.TotalDocks = v.TotalDocks
		stationResp = append(stationResp, stationlist)
	}

	//Check the API contain page number
	param := req.URL.Query().Get("page")
	if param != "" {
		page, _ := strconv.Atoi(param)
		start, end := Paginate((page - 1), 20, len(stationResp))
		pagedSlice := stationResp[start:end]
		sendResp, _ := json.Marshal(pagedSlice)
		w.Write(sendResp)
		return
	}
	sendResp, _ := json.Marshal(stationResp)
	w.Write(sendResp)
}

//Get station data - in Service
func GetStationsInService(w http.ResponseWriter, req *http.Request) {
	log.Info("Requested API - ", req.URL.Path)
	stationResp := CheckService(Cache.StationBeanLists, "In Service")
	sendResp, _ := json.Marshal(stationResp)
	w.Write(sendResp)
}

//Get station data - Not in Service
func GetStationsNotInService(w http.ResponseWriter, req *http.Request) {
	log.Info("Requested API - ", req.URL)
	stationResp := CheckService(Cache.StationBeanLists, "Not In Service")
	sendResp, _ := json.Marshal(stationResp)
	w.Write(sendResp)

}

//Get Station By Address
func GetStationByAddress(w http.ResponseWriter, req *http.Request) {
	log.Info("Requested API - ", req.URL.Path)
	vars := mux.Vars(req)
	station := vars["searchstring"]

	var stationLists []StationBeanList
	for _, service := range Cache.StationBeanLists {
		if strings.ToLower(service.StationName) == strings.ToLower(station) || strings.ToLower(service.StAddress1) == strings.ToLower(station) {
			stationLists = append(stationLists, service)
		}
	}
	if len(stationLists) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		Msg := UpdateErrorMessage("Info", "Data not found", 0, station, req.URL.Path)
		data, _ := json.Marshal(Msg)
		w.Write(data)
		return

	}
	sendResp, _ := json.Marshal(stationLists)
	w.Write(sendResp)
}

func FindDockable(w http.ResponseWriter, req *http.Request) {
	log.Info("Requested API - ", req.URL.Path)
	vars := mux.Vars(req)
	stationID, err := strconv.Atoi(vars["stationid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		Msg := UpdateErrorMessage("Info", "Station ID should be int", stationID, "", req.URL.Path)
		data, _ := json.Marshal(Msg)
		w.Write(data)
		return
	}
	bikeStoreTurn, err := strconv.Atoi(vars["bikestoreturn"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		Msg := UpdateErrorMessage("Info", "bikestoreturn should be string", stationID, "", req.URL.Path)
		data, _ := json.Marshal(Msg)
		w.Write(data)
		return
	}

	var resp Response
	var StationExist bool
	for _, service := range Cache.StationBeanLists {
		if service.ID == int64(stationID) {
			StationExist = true
			if service.AvailableDocks >= int64(bikeStoreTurn) {
				resp.Dockable = true
				resp.Message = "The requested space is available to Dock"
			} else {
				resp.Dockable = false
				resp.Message = "The requested space is not available to Dock"
			}
		}
	}
	if !StationExist {
		w.WriteHeader(http.StatusBadRequest)
		Msg := UpdateErrorMessage("Info", "Requested Station ID not found", stationID, "", req.URL.Path)
		data, _ := json.Marshal(Msg)
		w.Write(data)
		return

	}
	sendResp, _ := json.Marshal(resp)
	w.Write(sendResp)
}
