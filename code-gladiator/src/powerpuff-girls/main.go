package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	input := make([]string, 3)
	reader := bufio.NewReader(os.Stdin)

	input[0], _ = reader.ReadString('\n')
	input[1], _ = reader.ReadString('\n')
	input[2], _ = reader.ReadString('\n')

	noOfIngredients, _ := strconv.Atoi(strings.TrimSpace(input[0]))

	stringArr := strings.Split(strings.TrimSpace(input[1]), " ")

	quantityOfIngred := make([]int64, noOfIngredients)
	for i, v := range stringArr {
		val, _ := strconv.ParseInt(v, 10, 64)
		quantityOfIngred[i] = val
	}

	stringArr = strings.Split(strings.TrimSpace(input[2]), " ")

	quantityInLab := make([]int64, noOfIngredients)
	for i, v := range stringArr {
		val, _ := strconv.ParseInt(v, 10, 64)
		quantityInLab[i] = val
	}

	if noOfIngredients == 0 {
		fmt.Println(0)
		return

	}

	count := 0
	result := int64(math.MaxInt64)
	for count < noOfIngredients {

		if quantityOfIngred[count] > 0 {
			temp := quantityInLab[count] / quantityOfIngred[count]
			if result > temp {
				result = temp
			}
		}
		count++
	}
	if result == 2147483647 {
		result = 0
	}
	fmt.Println(result)
}
