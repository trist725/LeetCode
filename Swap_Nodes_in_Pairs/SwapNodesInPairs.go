package Swap_Nodes_in_Pairs

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 24
// https://leetcode.cn/problems/swap-nodes-in-pairs/

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	var (
		dummy = &ListNode{Next: head}
		cur   = dummy
	)

	for cur.Next != nil && cur.Next.Next != nil {
		n1 := cur.Next
		n2 := n1.Next
		n3 := n2.Next

		cur.Next = n2
		n2.Next = n1
		n1.Next = n3

		cur = n1
	}

	return dummy.Next
}
