package Intersection_of_Two_Linked_Lists_LCCI

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// https://leetcode.cn/problems/intersection-of-two-linked-lists-lcci/

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var (
		curA = headA
		curB = headB
	)

	if curA == nil || curB == nil {
		return nil
	}

	for curA != curB {
		if curA == nil {
			curA = headB
		} else {
			curA = curA.Next
		}
		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}

	return curA
}
