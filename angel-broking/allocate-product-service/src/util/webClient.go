package util

import (
	"io"
	"log"
	"net/http"
	"time"
)

var httpclient *http.Client

func init() {
	log.Println("Initializing http client..!!")
	httpclient = &http.Client{
		Timeout: time.Millisecond * 5000,
	}
}

// PerformRequest perform web request
func PerformRequest(method, requestPath string, data io.Reader) (*http.Response, error) {

	req, err := http.NewRequest(method, requestPath, data)
	if err != nil {
		return nil, err
	}
	return httpclient.Do(req)
}
