package main

import "fmt"

func main() {
	arr := []int{5, 7, 4, 10, 2, 8, 1}
	heapSort(arr)
}

type heapNode struct {
	Data  int
	left  *heapNode
	right *heapNode
}

func heapSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	heap := createHeap(arr)
	fmt.Println(heap)

}

func createHeap(arr []int) []int {

	heap := make([]int, len(arr))

	i := 0
	for i < len(arr) {
		heap[i] = arr[i]
		l := i
		k := (i + 1) / 2
		if heap[l] > heap[k] {

			for k > 0 {

				if heap[l] > heap[k-1] {
					heap[l], heap[k-1] = heap[k-1], heap[l]
				}
				l = k
				k = k / 2
			}
		}

		i++
	}

	return heap
}
