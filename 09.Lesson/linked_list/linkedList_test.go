package linked_list

import (
	"testing"
)

func TestGenerateList(t *testing.T) {
	head := generateList(3)

	var want = []int{1, 2, 3}
	for _, v := range want {
		if v != head.Val {
			t.Errorf("got %v want %v", head.Val, v)
		}
		if head.Next != nil {
			head = *head.Next
		}
	}
}

func TestReverseList(t *testing.T) {
	head := generateList(3)

	want := []int{3, 2, 1}
	var got []int
	newHead := reverseList(&head)
	for newHead != nil {
		got = append(got, newHead.Val)
		newHead = newHead.Next
	}

	for i, v := range want {
		if v != got[i] {
			t.Errorf("got %v want %v", got[i], v)
		}
	}
}

func TestReverseListRecursion(t *testing.T) {
	head := generateList(4)

	want := []int{4, 3, 2, 1}
	var got []int
	newHead := reverseList(&head)
	for newHead != nil {
		got = append(got, newHead.Val)
		newHead = newHead.Next
	}

	for i, v := range want {
		if v != got[i] {
			t.Errorf("got %v want %v", got[i], v)
		}
	}
}
