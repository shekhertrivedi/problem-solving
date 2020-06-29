package main

import "fmt"

type YearWisePopulation struct {
	Year  int
	Born  int
	Death int
	Max   int
}

func main() {
	y1 := YearWisePopulation{Year: 1, Born: 10, Death: 5}
	y2 := YearWisePopulation{Year: 2, Born: 10, Death: 5}
	y3 := YearWisePopulation{Year: 3, Born: 10, Death: 5}
	y4 := YearWisePopulation{Year: 4, Born: 10, Death: 5}
	y5 := YearWisePopulation{Year: 5, Born: 10, Death: 5}

	listOfYearsWisepop := []YearWisePopulation{y1, y2, y3, y4, y5}

	maxpop := getMaxPopulation(listOfYearsWisepop)

	fmt.Println(maxpop)

}

func getMaxPopulation(listOfYearsWisepop []YearWisePopulation) int {

	maxArr := make([]int, len(listOfYearsWisepop))
	maxNumber := 0
	maxArr[0] = listOfYearsWisepop[0].Born - listOfYearsWisepop[0].Death

	i := 1
	for i < len(listOfYearsWisepop) {
		maxArr[i] = maxArr[i-1] + listOfYearsWisepop[i].Born - listOfYearsWisepop[i].Death
		if maxArr[i] > maxNumber {
			maxNumber = maxArr[i]
		}
		i++
	}

	return maxNumber
}
