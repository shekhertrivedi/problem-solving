package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// fifth := ListNode{Val: 5, Next: nil}

	// fourth := ListNode{Val: 4, Next: &fifth}

	third := ListNode{Val: 3, Next: nil}

	second := ListNode{Val: 2, Next: &third}

	head := ListNode{Val: 1, Next: &second}

	// reverse := reverseList(&head)

	// for reverse != nil {
	// 	fmt.Println(reverse.Val)
	// 	reverse = reverse.Next
	// }

	reverse := reverseBetween(&head, 1, 2)

	for reverse != nil {
		fmt.Println(reverse.Val)
		reverse = reverse.Next
	}

}
func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	temp2 := head.Next
	head.Next = nil
	temp1 := head

	for temp2 != nil {
		temp := temp2.Next
		temp2.Next = temp1

		temp1 = temp2
		temp2 = temp
	}

	return temp1
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	origin := head
	count := 2
	for count < m {
		origin = origin.Next
		count++
	}

	index := 0

	temp := origin.Next
	for index < n-m && temp.Next != nil && temp.Next.Next != nil {

		temp1 := temp.Next

		temp2 := origin.Next

		origin.Next = temp1

		temp.Next = temp1.Next

		temp1.Next = temp2

		index++
	}

	return head
}
