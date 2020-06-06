package main

func main() {

}
// 4 5 6 7 8 1 2 3

func searchValue(arr []int, a int) {
	mid := len(arr) / 2
	start := 0
	end := len(arr)



	if arr[mid] == a {
		return mid
	} else if arr[mid] < arr[end] {

		

		} else {
			binarySearch(start, mid-1, arr)
		}

		
	} 



	binarySearch(start, mid, arr)
	binarySearch(mid+1, end, arr)
}

func binarySearch(start, end int, arr []int) int {

}
