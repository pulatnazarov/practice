package leetcode

import "fmt"

func DeleteDuplicates() {
	//delete duplicates from sorted linked list
	fmt.Println(deleteDuplicates(&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}))
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	curr := head
	for curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}
