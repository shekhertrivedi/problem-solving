package main

import (
	"reflect"
	"testing"
)

func Test_Paginate(t *testing.T) {

	tests := []struct {
		pageNumber    int
		pageSize      int
		sliceLength   int
		expectedStart int
		expectedEnd   int
	}{
		{
			pageNumber:    3,
			pageSize:      1,
			sliceLength:   12,
			expectedStart: 3,
			expectedEnd:   4,
		},
		{
			pageNumber:    1,
			pageSize:      3,
			sliceLength:   2,
			expectedStart: 2,
			expectedEnd:   2,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {

			start, end := Paginate(tt.pageNumber, tt.pageSize, tt.sliceLength)
			if start != tt.expectedStart || end != tt.expectedEnd {
				t.Errorf("Expected and Actual is not matching")
			}
		})
	}
}

func Test_CheckService(t *testing.T) {

	tests := []struct {
		station             []StationBeanList
		status              string
		expectedstationList []StationBeanList
	}{
		{
			station:             []StationBeanList{{ID: 304, StatusValue: "In Service"}},
			status:              "In Service",
			expectedstationList: []StationBeanList{{ID: 304, StatusValue: "In Service"}},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {

			list := CheckService(tt.station, tt.status)
			if !reflect.DeepEqual(list, tt.expectedstationList) {
				t.Errorf("Expected and Actual is not matching")
			}
		})
	}
}
