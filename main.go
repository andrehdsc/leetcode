package main

import "fmt"

func main() {
	mainCopy()

	ll := &LinkedList{
		Head: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	reverseList(ll)

	for i := ll.Head; i != nil; i = i.Next {
		fmt.Println(i.Val)
	}
}
