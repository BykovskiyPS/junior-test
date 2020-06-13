package linked_list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	var temp *ListNode

	for head != nil {
		temp = head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}

func reverseListRecursion(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	rv := reverseListRecursion(head.Next)
	head.Next.Next = head
	head.Next = nil
	return rv
}

func generateList(number int) ListNode {
	head := ListNode{Val: 1}
	point := &head
	for i := 2; i <= number; i++ {
		node := ListNode{Val: i}
		point.Next = &node
		point = point.Next
	}
	point.Next = nil
	return head
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
