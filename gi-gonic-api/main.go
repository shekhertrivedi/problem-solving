package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("start")
	r := getRouter()
	go func() {
		r.Run(":8051")
	}()

	// graceful shutdown
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("graceful shutdown")
}

func getRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", HomePage)
	r.POST("/", PostHomePage)
	r.GET("/query", QueryString)
	r.GET("/path/:name/:age", PathVariable)
	r.POST("/body", PostBody)

	return r
}

func HomePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "success GET",
	})
}

func PostHomePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "success POST",
	})
}

func QueryString(c *gin.Context) {
	name, ok := c.GetQuery("name")
	age, ok := c.GetQuery("age")

	fmt.Println(ok)

	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}

func PathVariable(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}

type Person struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func PostBody(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}

	p := &Person{}
	err = json.Unmarshal(value, p)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(p)
	c.JSON(http.StatusOK, p)
}
