package main

import (
	"bufio"
	"fmt"
	"os"
)

func main123() {
	reader := bufio.NewReader(os.Stdin)
	input := make([]rune, 0)
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println(err)
		}
		if char == '0' {
			break
		}
		input = append(input, char)

	}
	boringDays := 0
	for i := 1; i < len(input); i++ {
		if input[i] == input[i-1] {
			boringDays++
		}
	}
	fmt.Println(boringDays)
}
