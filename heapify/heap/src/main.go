package main

import "fmt"

func main() {

	arr := []int{16, 30, 8, 20, 50, 10}
	arr = createHeap(arr)
	sortHeap(arr)

}

func sortHeap(arr []int) {
	i := 0
	for i < len(arr) {

		element := arr[0]
		arr[i] = arr[len(arr)-1-i]

		adjustHeapSort(arr, 0, len(arr)-i)

		arr[len(arr)-1-i] = element
		i++
	}
	fmt.Println(arr)
}

func adjustHeapSort(arr []int, i, limit int) {

	left := (2 * (i + 1)) - 1
	right := (2*(i+1) + 1) - 1

	if left < limit && right < limit {
		index := max(arr, left, right)
		if arr[i] < arr[index] {
			arr[i], arr[index] = arr[index], arr[i]
			adjustHeapSort(arr, index, limit)
		}
	}
	return
}

func max(arr []int, a, b int) int {
	if arr[a] > arr[b] {
		return a
	}
	return b
}

func createHeap(arr []int) []int {
	heap := make([]int, len(arr))
	heap[0] = arr[0]
	for i := 1; i < len(arr); i++ {

		heap[i] = arr[i]

		adjustHeap(heap, i)
	}
	return heap
}

func adjustHeap(heap []int, i int) {

	index := (i+1)/2 - 1
	if index != -1 && heap[index] < heap[i] {
		heap[index], heap[i] = heap[i], heap[index]
		adjustHeap(heap, index)
	}
	return
}
