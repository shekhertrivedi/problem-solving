package main

import (
	"fmt"
	"strings"
)

var res string

func main() {
	//reverseString("abcd")

	s1 := "acd"
	s2 := "acd"

	fmt.Println(strings.Compare(s1, s2))

	s1 = "acb"
	s2 = "acd"

	fmt.Println(strings.Compare(s1, s2))

	s1 = "acd"
	s2 = "acb"
	fmt.Println(strings.Compare(s1, s2))

}

func reverseString(s string) {

	if len(s) == 0 {
		return
	}
	res := make([]rune, 0)
	sArr := []rune(s)
	i := len(s) - 1
	for i >= 0 {
		res = append(res, sArr[i])
		i = i - 1
	}
	fmt.Println(string(res))
}
