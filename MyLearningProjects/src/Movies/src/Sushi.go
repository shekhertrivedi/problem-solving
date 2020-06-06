package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

/*

Question is not very clear, As there could be multiple solutuion.
In current implementation breaking on first solution

*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	input := make([]int, 0)
	for scanner.Scan() {
		command := scanner.Text()
		num, _ := strconv.Atoi(command)
		if num == 0 {
			break
		}
		input = append(input, num)
	}

	for _, v := range input {

		if v < 0 {
			fmt.Println("No solution")
			continue
		}

		sqRoot := int(math.Ceil(math.Sqrt(float64(v))))
		var isPrime bool

		i := 2
		for i < sqRoot {
			if (v % i) == 0 {
				isPrime = false
				break
			}
			i++
		}
		if isPrime {
			fmt.Println("No solution")
			continue
		}
		sol := (v / i) - 1
		if sol == 1 {
			fmt.Println("No solution")
			continue
		}
		fmt.Println(sol) // fmt.Println(i-1)
	}
}
