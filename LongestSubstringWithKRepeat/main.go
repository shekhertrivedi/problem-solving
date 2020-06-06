package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := "ababbc"
	fmt.Println(longestSubstring(s, 2))
}

func stringProcessing(s string) {

	//print ascii for chars
	fmt.Println(s[0])
	fmt.Println(reflect.TypeOf(s[0]))

	// print char of string
	fmt.Println(string(s[1]))

	// to chat array

	arr := []rune(s)
	fmt.Println(arr[0])
	fmt.Println(string(arr[0]))

	// making changes to the existing string
	arr[0] = 'z'

	//convert char array to string
	newString := string(arr)
	fmt.Println(newString)
}

func longestSubstring(s string, k int) int {

	m := make(map[rune]int)
	for _, v := range s {
		if x, ok := m[v]; ok {
			m[v] = x + 1
		} else {
			m[v] = 1
		}
	}

	var possibleResults []string
	resultString := ""

	for i := 0; i < len(s); i++ {

		if v, _ := m[rune(s[i])]; v < k {

			if ifValid(resultString, k) {
				possibleResults = append(possibleResults, resultString)
				resultString = ""
				continue
			}

		}

		resultString = resultString + string(s[i])
	}

	if ifValid(resultString, k) {
		possibleResults = append(possibleResults, resultString)
	}
	max := 0
	for _, v := range possibleResults {
		if len(v) > max {
			max = len(v)
		}
	}
	return max
}

func ifValid(s string, k int) bool {

	// check if the string has all the char k or more times in the string
	m := make(map[rune]int)
	for _, v := range s {
		if x, ok := m[v]; ok {
			m[v] = x + 1
		} else {
			m[v] = 1
		}
	}

	for i := 0; i < len(s); i++ {

		if v, _ := m[rune(s[i])]; v < k {
			return false
		}

	}
	return true
}
