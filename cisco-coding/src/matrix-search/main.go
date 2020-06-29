package main

/*

OSI Layers

1. Application (Http,FTP,telnet)
PResentation (SHA256, SHA128,)
Session ()
Transport (ELB)
Network
DAta
Physical 1

TCP/IP


1  2  3  4  5
6  7  8  9  10
11 12 13 14 15
16 17 18 19 20

*/

import "fmt"

func main() {

}
func getnumber(arr [][]int, row, column, target int) bool {

	// identify the row
	start := 0
	end := row - 1
	mid := (start + end) / 2

	rowTobeLooked := 0
	for {

		if (end - start) == 1 {
			rowTobeLooked = mid
			break
		}

		if arr[mid][0] == target {
			fmt.Println(mid, 0)
			return true
		} else if arr[mid][0] < target {
			start = mid //2
			//end = end //3
			mid = (start + end) / 2 //2
		} else {
			start = mid
			mid = (start + end) / 2
		}

	}

	start = 0
	end = column - 1
	mid = (start + end) / 2
	for end > start {
		if arr[rowTobeLooked][mid] == target {
			fmt.Println(mid, 0)
			return true
		} else if arr[rowTobeLooked][mid] < target {
			start = mid
			mid = (start + end) / 2
		} else {
			start = mid
			mid = (start + end) / 2
		}
	}

	return false
}
