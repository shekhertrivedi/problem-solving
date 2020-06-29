package main

import (
	_ "problem-solving/angel-broking/add-product-service/src/repository"
	"problem-solving/angel-broking/add-product-service/src/util"
	"problem-solving/angel-broking/add-product-service/src/web"
)

func main() {
	util.ConfigCB()
	web.NewRouter()
}
