package main


type ListNode struct {
    Val int
    Next *ListNode
}

type LinkedList struct {
	Head *ListNode
}
 
func reverseList(ll *LinkedList) {
    curr := ll.Head
    var prev *ListNode
    var next *ListNode

    for curr != nil {
        next = curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    ll.Head = prev
}