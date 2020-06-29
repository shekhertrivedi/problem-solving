package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
)

func init() {
	os.Setenv("LOGLEVELS", "VERBOSE")
	log = NewLogger("NBC")
	stationlist := StationBeanList{}
	stationlist.ID = 304
	stationlist.StationName = "Broadway & Battery Pl"
	stationlist.AvailableDocks = 16
	stationlist.AvailableBikes = 15
	stationlist.TotalDocks = 33
	stationlist.StatusValue = "In Service"
	stationlist.StAddress1 = "Broadway & Battery Pl"
	stationlist.StAddress2 = "Broadway"
	stationlist.City = "mockcity"
	stationlist.PostalCode = "0000"

	Cache = &Station{ExecutionTime: "2020-03-16 10:28:01 AM", StationBeanLists: []StationBeanList{stationlist, stationlist}}

}

//Test Get Station list
func TestGetStations(t *testing.T) {

	stationlist := StationBeanList{}
	stationlist.StationName = "Broadway & Battery Pl"
	stationlist.StAddress1 = "Broadway & Battery Pl"
	stationlist.StAddress2 = "Broadway"
	stationlist.City = "mockcity"
	stationlist.PostalCode = "0000"
	stationlist.AvailableBikes = 15
	stationlist.TotalDocks = 33
	expected := []StationBeanList{stationlist, stationlist}

	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "http://localhost/Station", nil)
	if err != nil {
		fmt.Errorf("Cannot read the request")
	}
	// test station list
	GetStations(w, r)
	if w.Code != http.StatusOK {
		t.Error("Expected 200 but got", w.Code)
	}
	respByte, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Error("fail to read response data")
		return
	}
	var responseData []StationBeanList
	json.Unmarshal(respByte, &responseData)
	//test response data
	if !reflect.DeepEqual(responseData, expected) {
		t.Errorf("Expected %v but got %v", expected, responseData)
	}

	///test station with page
	q := r.URL.Query()
	q.Add("page", "1")
	r.URL.RawQuery = q.Encode()
	GetStations(w, r)
	if w.Code != http.StatusOK {
		t.Error("Expected 200 but got", w.Code)
	}

}

//Test Station in service
func TestGetStationsInService(t *testing.T) {
	stationlist := StationBeanList{}
	stationlist.ID = 304
	stationlist.AvailableDocks = 16
	stationlist.StationName = "Broadway & Battery Pl"
	stationlist.StAddress1 = "Broadway & Battery Pl"
	stationlist.StAddress2 = "Broadway"
	stationlist.City = "mockcity"
	stationlist.PostalCode = "0000"
	stationlist.AvailableBikes = 15
	stationlist.TotalDocks = 33
	stationlist.StatusValue = "In Service"

	expected := []StationBeanList{stationlist, stationlist}

	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "http://localhost/station/service-in", nil)
	if err != nil {
		fmt.Errorf("Cannot read the request")
	}
	GetStationsInService(w, r)
	if w.Code != http.StatusOK {
		t.Error("Expected 200 but got", w.Code)
	}
	respByte, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Error("fail to read response data")
		return
	}
	var responseData []StationBeanList
	json.Unmarshal(respByte, &responseData)
	if !reflect.DeepEqual(responseData, expected) {
		t.Errorf("Expected %v but got %v", expected, responseData)
	}
}

func TestGetStationsNotInService(t *testing.T) {
	stationlist := StationBeanList{}
	stationlist.ID = 304
	stationlist.AvailableDocks = 16
	stationlist.StationName = "Broadway & Battery Pl"
	stationlist.StAddress1 = "Broadway & Battery Pl"
	stationlist.StAddress2 = "Broadway"
	stationlist.City = "mockcity"
	stationlist.PostalCode = "0000"
	stationlist.AvailableBikes = 15
	stationlist.TotalDocks = 33
	stationlist.StatusValue = "Not In Service"
	expected := []StationBeanList{stationlist}
	Cache.StationBeanLists = append(Cache.StationBeanLists, stationlist)
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "http://localhost/station/notinservice", nil)
	if err != nil {
		fmt.Errorf("Cannot read the request")
	}
	GetStationsNotInService(w, r)
	if w.Code != http.StatusOK {
		t.Error("Expected 200 but got", w.Code)
	}

	respByte, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Error("fail to read response data")
		return
	}
	var responseData []StationBeanList
	json.Unmarshal(respByte, &responseData)

	if !reflect.DeepEqual(responseData, expected) {
		t.Errorf("Expected %v but got %v", expected, responseData)
	}

}

func TestGetStationByAddress(t *testing.T) {

	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "http://localhost/station/Broadway & Battery Pl", nil)
	if err != nil {
		fmt.Errorf("Cannot read the request")
	}
	///test with valid data
	r = mux.SetURLVars(r, map[string]string{"searchstring": "Broadway & Battery Pl"})
	GetStationByAddress(w, r)
	if w.Code != http.StatusOK {
		t.Error("Expected 200 but got", w.Code)
	}
	///Test with invalid data
	r = mux.SetURLVars(r, map[string]string{"searchstring": "invalid data"})
	GetStationByAddress(w, r)
	if w.Code != http.StatusOK {
		t.Error("Expected 200 but got", w.Code)
	}
	//invalid data
}
func TestFindDockable(t *testing.T) {

	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "http://localhost/station/id/bikestore", nil)
	if err != nil {
		fmt.Errorf("Cannot read the request")
	}
	//Valid test
	r = mux.SetURLVars(r, map[string]string{"stationid": "304", "bikestoreturn": "2"})
	FindDockable(w, r)
	if w.Code != http.StatusOK {
		t.Error("Expected 200 but got", w.Code)
	}
	//invalid bikestoreturn
	r = mux.SetURLVars(r, map[string]string{"stationid": "304", "bikestoreturn": "100"})
	FindDockable(w, r)
	if w.Code != http.StatusOK {
		t.Error("Expected 200 but got", w.Code)
	}
	//invalid station id
	r = mux.SetURLVars(r, map[string]string{"stationid": "100000123", "bikestoreturn": "100"})
	FindDockable(w, r)
	if w.Code != http.StatusOK {
		t.Error("Expected 200 but got", w.Code)
	}
	//stationID not int
	r = mux.SetURLVars(r, map[string]string{"stationid": "sampleid", "bikestoreturn": "100"})
	FindDockable(w, r)
	if w.Code != http.StatusOK {
		t.Error("Expected 200 but got", w.Code)
	}
	//bikestoreturn not int
	r = mux.SetURLVars(r, map[string]string{"stationid": "304", "bikestoreturn": "samplebike"})
	FindDockable(w, r)
	if w.Code != http.StatusOK {
		t.Error("Expected 200 but got", w.Code)
	}
}

func TestMiddleware(t *testing.T) {

	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "http://localhost/station/id/bikestore", nil)
	if err != nil {
		fmt.Errorf("Cannot read the request")
	}
	Middleware(w, r, GetStations)
	if w.Code != http.StatusOK {
		t.Error("Expected 200 but got", w.Code)
	}
	//Cache nil with invalid url
	Cache = nil
	URL = "http://localhost:4000/middleware"
	Middleware(w, r, GetStations)
	if w.Code != http.StatusOK {
		t.Error("Expected 500 but got", w.Code)
	}
	//Cache nil with mockurl url
	//mock server
	Cache = nil
	mockts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

	}))
	defer mockts.Close()
	URL = mockts.URL
	Middleware(w, r, GetStations)
	if w.Code != http.StatusOK {
		t.Error("Expected 500 but got", w.Code)
	}

}
