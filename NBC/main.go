package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

var Cache *Station
var log Logger
var URL string = "https://feeds.citibikenyc.com/stations/stations.json"

func Middleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	if Cache == nil {
		log.Info("No Data found in Cache ...Loading")
		//Load Data to cache
		client := http.Client{
			Timeout: time.Duration(5 * time.Second),
		}

		response, err := client.Get(URL)
		if err != nil || response == nil {
			w.WriteHeader(http.StatusInternalServerError)
			Msg := UpdateErrorMessage("Error", "Unable to Hit the station URL with requested time", 0, "", r.URL.Path)
			data, _ := json.Marshal(Msg)
			w.Write(data)
			return
			//}
		}
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			Msg := UpdateErrorMessage("Error", "Unable to read response body", 0, "", r.URL.Path)
			data, _ := json.Marshal(Msg)
			w.Write(data)
			return
		}

		err = json.Unmarshal(data, &Cache)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			Msg := UpdateErrorMessage("Error", "Unable to Unmarshal station Json data", 0, "", r.URL.Path)
			data, _ := json.Marshal(Msg)
			w.Write(data)
			return
		}

	} else {
		log.Info("Data already Exist in Cache")
	}
	next(w, r)
}

func main() {
	os.Setenv("LOGLEVELS", "VERBOSE")
	log = NewLogger("NBC")
	log.Info("Application started...")
	router := mux.NewRouter()

	router.HandleFunc("/stations", GetStations)
	router.HandleFunc("/stations/in-service", GetStationsInService)
	router.HandleFunc("/stations/not-in-service", GetStationsNotInService)
	router.HandleFunc("/stations/{searchstring}", GetStationByAddress)
	router.HandleFunc("/dockable/{stationid}/{bikestoreturn}", FindDockable)

	n := negroni.New(negroni.HandlerFunc(Middleware))
	n.UseHandler(router)
	n.Run(":4000")
}
