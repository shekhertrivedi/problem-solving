package main

import "fmt"

func main() {
	//nums := []int{3, 4, 2, 3}
	//nums := []int{1, 5, 4, 6, 7, 10, 8, 9}
	nums := []int{2, 3, 3, 2, 4}
	fmt.Println(checkPossibility(nums))
}

/*

Given an array nums with n integers, your task is to check if it could become non-decreasing by modifying at most 1 element.

We define an array is non-decreasing if nums[i] <= nums[i + 1] holds for every i (0-based) such that (0 <= i <= n - 2).



Example 1:

Input: nums = [4,2,3]
Output: true
Explanation: You could modify the first 4 to 1 to get a non-decreasing array.
Example 2:

Input: nums = [4,2,1]
Output: false
Explanation: You can't get a non-decreasing array by modify at most one element.


Constraints:

1 <= n <= 10 ^ 4
- 10 ^ 5 <= nums[i] <= 10 ^ 5

*/

func checkPossibility(nums []int) bool {
	changesReq := 0
	i := 0
	for i < len(nums)-1 {
		if nums[i] > nums[i+1] {
			changesReq++
			if changesReq > 1 {
				return false
			}
			nums = placeElementAtCorrectPos(nums, i)
			i = 0
			continue
		}
		i++
	}
	return true
}

func placeElementAtCorrectPos(nums []int, i int) []int {
	m := i
	for m < len(nums)-1 {
		if nums[m] > nums[m+1] {
			nums[m], nums[m+1] = nums[m+1], nums[m]
		} else {
			break
		}
		m++
	}
	return nums
}
