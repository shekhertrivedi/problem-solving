package main

import (
	"fmt"
)

//6,7,8,5,9,3,5,4

func main() {
	fmt.Println("Hello, playground")
	//arr := []int{6, 7, 8, 5, 9, 3, 5, 4}
	arr := []int{3, 4, 5, 5, 6, 7, 8, 9}
	removeDuplicates(arr)
}

// assumption sorted array
// 3,4,5,5,6,7,8,9
func removeDuplicates(arr []int) {
	result := make([]int, 0)
	i := 0
	k := 0
	for i < len(arr)-1 {

		if arr[i] == arr[i+1] {
			arr[k] = arr[i]
			k++
			result = append(result, arr[i])
			//duplicateStart := i + 1
			//duplicateElement := arr[i+1]

			j := i + 1
			for j < len(arr) {
				if arr[j] == arr[i+1] {
					j++
					continue
				}
				break
			}
			result = append(result, arr[j])
			arr[k] = arr[j]
			k++
			i = j + 1
		} else {
			result = append(result, arr[i])
			arr[k] = arr[i]
			k++
			i++
		}
	}

	//for _, v := range result {
	//	fmt.Println(v)
	//}
	m := 0
	for m < k {
		fmt.Println(arr[m])
		m++
	}

}
