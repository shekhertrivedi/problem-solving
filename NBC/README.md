# NBC News Digital: Go Take-home Coding Challenge
This is an backend applictaion. This service will handle below requests

- GetStation
- GetStationByInService
- GetStationByNotInService
- GetStationBysearchstring
- FindDockabeUsingAvailableBikes


GetStation - This will fetch all the station details along with pagenite option
GetStationByInService - This will provide on In service station details
GetStationByNotInService - This will provide on Not In service station details
GetStationBysearchstring - This will provide search related station Details without case sensitive
FindDockabeUsingAvailableBikes - This will find and return requested docks available or not.

To ENABLE LOG :
- SET LOGLEVEL=VERBOSE(for now hardcoded)

URL USED TO FETCH JSON PAYLOAD :
- https://feeds.citibikenyc.com/stations/stations.json

APPLICATION RUNS ON :
- port 4000

DEPENDENCY PACKAGE :

- go get  "github.com/gorilla/mux"
- go get  "github.com/urfave/negroni"

TO BUILD :

- CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build 

To RUN :
 ./NBC

TO RUN TEST CASE :

- go test --coverprofile=cover.out

To SEE COVERAGE REPORT

- go tool cover -html=cover.out

