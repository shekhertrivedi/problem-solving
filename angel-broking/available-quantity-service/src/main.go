package main

import (
	"problem-solving/angel-broking/available-quantity-service/src/util"
	"problem-solving/angel-broking/available-quantity-service/src/web"
)

func main() {
	util.ConfigCB()
	web.NewRouter()
}
