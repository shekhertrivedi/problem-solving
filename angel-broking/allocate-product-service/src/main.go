package main

import (
	"problem-solving/angel-broking/allocate-product-service/src/util"
	"problem-solving/angel-broking/allocate-product-service/src/web"
)

func main() {
	util.ConfigCB()
	web.NewRouter()
}
