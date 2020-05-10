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
	fifth := ListNode{Val: 5, Next: nil}

	fourth := ListNode{Val: 4, Next: &fifth}

	third := ListNode{Val: 3, Next: &fourth}

	second := ListNode{Val: 2, Next: &third}

	head := ListNode{Val: 1, Next: &second}

	reverse := reverseList(&head)

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
