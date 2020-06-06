/*
 *Read input from STDIN. Print your output to STDOUT
 *Use fmt.Scanf to read input from STDIN and fmt. Println to write output to STDOUT
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	noOfElements, _ := strconv.Atoi(strings.TrimSpace(input))

	if noOfElements < 1 {
		fmt.Print(-1)
		return
	}

	inputArr := make([]int, noOfElements)

	i := 0
	for i < noOfElements {
		element, _ := reader.ReadString('\n')
		inputArr[i], _ = strconv.Atoi(strings.TrimSpace(element))
		i++
	}

	k := noOfElements

	for i = k - 1; i >= 0; i-- {
		if inputArr[i] == k {
			k = k - 1
		}
	}
	fmt.Print(k)
}
