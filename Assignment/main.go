package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gojektech/heimdall/httpclient"
	"github.com/gorilla/mux"
)

type Incident struct {
	TaskID      string `json:"taskId,omitempty"`
	IncientID   string `json:"incidentId,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type Note struct {
	TaskID  string `json:"taskId,omitempty"`
	NoteId  string `json:"noteId,omitempty"`
	Details string `json:"details,omitempty"`
}

type Contact struct {
	TaskID         string `json:"taskId,omitempty"`
	ContactID      string `json:"contactId,omitempty"`
	ContactDetails string `json:"condetails,omitempty"`
}

func main() {
	// fmt.Println("Inside main")
	// http.ListenAndServe(":8050", ConfigureRouter())

	c1 := Contact{TaskID: "123"}
	b, _ := json.Marshal(c1)

	m := make(map[string]interface{})
	json.Unmarshal(b, &m)
	fmt.Println(m)
}

func ConfigureRouter() *mux.Router {
	r := mux.NewRouter()
	//r.HandleFunc("/", jokeHandler)
	r.HandleFunc("/jsonEx", jsonHandler).Methods(http.MethodPost)
	http.Handle("/", r)
	return r
}

func jsonHandler(resp http.ResponseWriter, r *http.Request) {
	i := Incident{}

	fmt.Println("Length of I: ")
	n := Note{}
	c := Contact{}

	data, err := ioutil.ReadAll(r.Body)

	t1 := bytes.NewReader(data)
	t2 := bytes.NewReader(data)
	t3 := bytes.NewReader(data)

	//t3 := r.Body
	err = json.NewDecoder(t3).Decode(&i)
	fmt.Println("Incident: ", i, err)
	err = json.NewDecoder(t1).Decode(&n)
	fmt.Println("Note: ", n, err)
	err = json.NewDecoder(t2).Decode(&c)
	fmt.Println("Contact: ", c, err)

	res, _ := json.Marshal(i)

	resp.WriteHeader(201)
	resp.Write(res)
}

func jokeHandler(resp http.ResponseWriter, r *http.Request) {

	p := Person{}
	err := Get("http://uinames.com/api/", &p)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Error occured while fetching name"))
	}

	jd := JokeDetails{}

	path := "http://api.icndb.com/jokes/random?firstName=%v&lastName=%v&limitTo=[nerdy]"
	url := fmt.Sprintf(path, "John", "Doe")

	finalName := p.Name + " " + p.Surname
	err = Get(url, &jd)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Error occured while fetching joke"))
	}

	joke := strings.Replace(jd.Value.Joke, "John Doe", finalName, 1)
	logInfo(fmt.Sprintf("First Name: %v Last Name: %v Joke: %v", p.Name, p.Surname, joke))

	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte(joke))

}

func Get(path string, model interface{}) error {

	timeout := 1000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

	res, err := client.Get(path, nil)
	if err != nil {
		return err
	}
	return json.NewDecoder(res.Body).Decode(model)
}

type Person struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Gender  string `json:"gender"`
	Region  string `json:"region"`
}

//{ "type": "success", "value": { "id": 487, "joke": "No statement can catch the JohnDoeException.", "categories": ["nerdy"] } }

type JokeDetails struct {
	Type  string           `json:"type"`
	Value JokeDetailsValue `json:"value"`
}

type JokeDetailsValue struct {
	ID         int      `json:"id"`
	Joke       string   `json:"joke"`
	Categories []string `json:"categories"`
}

var logInfo = func(msg string) {
	log.Printf(msg)
}
