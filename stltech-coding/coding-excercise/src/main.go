package main

import "fmt"

var X int

type Node struct {
	data int
	Next *Node
}

// 1 -> 2 -> 3 -> 4 -> 5 -> 6 ------ nil
//                      |     |
func main() {

	// n4 := Node{data: 5, Next: nil}
	// n3 := Node{data: 4, Next: &n4}
	// n2 := Node{data: 3, Next: &n3}
	// n1 := Node{data: 2, Next: &n2}

	head := Node{data: 1, Next: nil}
	// n4.Next = &head
	// n := &head
	// for n != nil {
	// 	fmt.Println(n.data)
	// 	n = n.Next
	// }
	s := &Stack{Head: &head}
	fmt.Println(s.peek())
	// 1----nil

	fmt.Println("POP")
	fmt.Println(s.pop()) // 1
	fmt.Println(s.Head)  // nil
	fmt.Println(s.peek())
	fmt.Println("PUSH")
	s.push(5)
	fmt.Println(s.Head) // non-nil 5

	fmt.Println("PEEK")
	fmt.Println(s.peek()) // 5

	//fmt.Println(findLoop(&head))
}

type Stack struct {
	Head *Node
}

/*
nil
1 --- nil
2 ---- 1 -- nil
*/

func (s *Stack) push(data int) {
	if s.Head == nil {
		head := &Node{data: data, Next: nil}
		s.Head = head
		return
	}

	s.Head = &Node{data: data, Next: s.Head}
}

func (s *Stack) pop() int {

	if s.Head == nil {
		return -1
	}
	retVal := s.Head.data
	if s.Head.Next == nil {
		s.Head = nil
	} else {
		s.Head = s.Head.Next
	}
	return retVal
}

func (s *Stack) peek() int {
	if s.Head == nil {
		return -1
	}
	return s.Head.data
}

func findLoop(head *Node) bool {

	slow := head
	fast := head

	for fast != nil {

		if fast.Next == nil || fast.Next.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}
