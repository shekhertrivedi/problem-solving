package main

import "fmt"

func main() {

	arr := [][]int{
		{1, 2, 3, 4, 5, 6},
		{4, 5, 6, 7, 8, 9},
		{10, 11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20, 21},
	}

	printMatrix(arr)
}

func printMatrix(arr [][]int) {
	columns := len(arr[0])
	rows := len(arr)

	fmt.Println(rows, columns)
	i := 0
	for i < rows {
		temp := i
		j := 0
		for temp >= 0 {
			fmt.Print(arr[temp][j])
			temp = temp - 1
			j++
		}
		fmt.Println()
		i++
	}

	j := 1
	i = rows - 1
	for j < columns {

		temp := j
		row := i
		for row < rows && temp < columns {
			fmt.Println(arr[row][temp])
			temp++
			row = row - 1
		}
		j++

	}
}
