package Reverse_Linked_List

// 206
// https://leetcode.cn/problems/reverse-linked-list/

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

func ReverseList_1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var (
		cur  = head
		prev *ListNode
		next = cur.Next
	)

	for cur != nil {
		cur.Next = prev
		prev = cur
		cur = next
		if next != nil {
			next = next.Next
		}
	}

	return prev
}

// 递归
func ReverseList_2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := ReverseList_2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
