package main

import "fmt"

func main() {
	s := "abc"
	getPermutation(s, 0)
}

func getPermutation(s string, count int) {
	if count == len(s) {
		fmt.Println(s)
		return
	}

}
