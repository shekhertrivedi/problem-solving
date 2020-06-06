package main

import "fmt"

func main() {
	permutations12("abc")

	//fmt.Println(swap("abcd", 1, 2))
}

func permutations12(s string) {

	if len(s) > 0 {
		permute(s, 0, len(s)-1)
	}
}

func permute(s string, start, end int) {
	if start == end {
		fmt.Println(s)
	} else {
		for i := start; i <= end; i++ {
			s = swap(s, start, i)
			permute(s, start+1, end)
			s = swap(s, start, i)
		}
	}
}

func swap(a string, i, j int) string {
	r := []rune(a)
	r[i], r[j] = r[j], r[i]

	return string(r)
}

func myPermute(s string, start, end int) {
	if start == end {
		fmt.Println(s)
	}
}
