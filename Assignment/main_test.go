package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConfigureRouter(t *testing.T) {
	r := ConfigureRouter()

	if r == nil {
		t.Errorf("Failed to configure the router")
	}

}

//Test getNamejokeHandler - positive test case
func Test_getJokeHandler(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	jokeHandler(response, request)
	if response.Code != http.StatusOK {
		t.Fatalf("Non-expected status code%v:\n\tbody: %v", "200", response.Code)
	}
}

func TestGet(t *testing.T) {
	p := Person{}
	path := "http://uinames.com/api/"
	err := Get(path, &p)
	if err != nil {
		t.Errorf(fmt.Sprintf("Failed while making api call for: %v", path))
	}
}

//TODO : need to write negative test case
