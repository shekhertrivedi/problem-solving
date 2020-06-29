package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ContinuumLLC/rmm-its-api/src/listener"
	"github.com/ContinuumLLC/rmm-its-api/src/local/logging"
	platConfig "github.com/ContinuumLLC/rmm-its-api/src/local/platform/config"
	"github.com/ContinuumLLC/rmm-its-api/src/web"

	"github.com/ContinuumLLC/rmm-its-api/src/config"
	"github.com/ContinuumLLC/rmm-its-api/src/messaging"
	l "github.com/sirupsen/logrus"
)

const (
	serverConfigName = "ServerConfig"
	serverConfigExt  = "json"
)

var (
	configMgrs  platConfig.Configurations
	webListener web.Listener
	// Build represent git build no
	Build string
)

type middleware func(next http.HandlerFunc) http.HandlerFunc

type User struct {
	UserName string `json:"abc,omitempty"`
	Password string
	Userid   int
}

func main() {
	//viper.SetConfigName("config") // name of config file (without extension)
	//lt := chainMiddleware(withLogging, withTracing)

	l.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	lt := chainMiddleware(checkAuthorization)
	http.Handle("/", lt(home))
	http.Handle("/about", lt(about))
	http.Handle("/about", lt(parseJson))
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}

}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("do something home")
	w.WriteHeader(http.StatusOK)
	msg := "{\"a\":\"x\"}"
	//jso, _ := json.Marshal(msg)
	w.Write([]byte(msg))

}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Println("do something about")

	user1 := User{"", "pwd", 4}
	jso, _ := json.Marshal(user1)
	w.Write(jso)
}

func parseJson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("do something about")

	user1 := User{"", "pwd", 4}
	//json.Unmarshal(r.GetBody, &user1)
	jso, _ := json.Marshal(user1)
	w.Write(jso)
}

func checkAuthorization(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Make API call to check authorization")
		next.ServeHTTP(w, r)
	}
}

// chainMiddleware provides syntactic sugar to create a new middleware
// which will be the result of chaining the ones received as parameters.
func chainMiddleware(mw ...middleware) middleware {
	return func(final http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			last := final
			for i := len(mw) - 1; i >= 0; i-- {
				last = mw[i](last)
			}
			last(w, r)
		}
	}
}

func withLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Logged connection from %s", r.RemoteAddr)
		next.ServeHTTP(w, r)
	}
}

func withTracing(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Tracing request for %s", r.RequestURI)
		next.ServeHTTP(w, r)
	}
}
