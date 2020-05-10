package main

import "fmt"

func main() {
	s := "ababbc"
	fmt.Println(longestSubstring(s, 2))
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
