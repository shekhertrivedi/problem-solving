package main

import "fmt"

func main() {
	//printFibonacci(4)
	//arr := []int{1, 2, 3, 4, 5, 6}
	arr := []int{1, 8, 3, 4, 5, 0}
	maxArr123(arr)
}

func maxArr123(arr []int) {
	first := arr[0]
	second := arr[0]
	for i, v := range arr {
		if i > 0 && v > second {
			if v > first {
				second = first
				first = v
			} else {
				second = v
			}
		}
	}
	fmt.Println(first, second)
}

// max no of array

func maxArr(arr []int) {
	max := arr[0]
	for i, v := range arr {
		if i > 0 && v > max {
			max = v
		}
	}
	fmt.Println(max)
}

//0,1,1,2,3,5,8

func printFibonacci(n int) {

	first := 0
	second := 1
	for n > 0 {
		fmt.Println(first)
		temp := first + second
		first = second
		second = temp
		n = n - 1
	}

}
