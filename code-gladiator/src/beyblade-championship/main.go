package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	noOfTestCases, _ := strconv.Atoi(strings.TrimSpace(input))

	output := make([]int, noOfTestCases)

	count := 0
	for count < noOfTestCases {

		testCase := make([]string, 3)

		testCase[0], _ = reader.ReadString('\n')
		testCase[1], _ = reader.ReadString('\n')
		testCase[2], _ = reader.ReadString('\n')

		noOfPlayers, _ := strconv.Atoi(strings.TrimSpace(testCase[0]))

		stringArr := strings.Split(strings.TrimSpace(testCase[1]), " ")

		gRevMembers := make([]int64, noOfPlayers)
		for i, v := range stringArr {
			val, _ := strconv.ParseInt(v, 10, 64)
			gRevMembers[i] = val
		}

		stringArr = strings.Split(strings.TrimSpace(testCase[2]), " ")

		opponentMembers := make([]int64, noOfPlayers)
		for i, v := range stringArr {
			val, _ := strconv.ParseInt(v, 10, 64)
			opponentMembers[i] = val
		}

		sort.Slice(gRevMembers, func(i, j int) bool { return gRevMembers[i] < gRevMembers[j] })

		sort.Slice(opponentMembers, func(i, j int) bool { return opponentMembers[i] < opponentMembers[j] })

		j := 0
		i := 0
		win := 0
		for i < noOfPlayers {
			if gRevMembers[i] > opponentMembers[j] {
				win++
				i++
				j++
				continue
			}
			i++
		}
		output[count] = win
		count++
	}
	for _, v := range output {
		fmt.Println(v)
	}
}
