package Remove_Nth_Node_From_End_of_List

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 19
// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1 <= n
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	var (
		dummy    = &ListNode{Next: head}
		slowPrev = dummy
		slow     = head
		fast     = slow
	)

	for i := 1; i < n; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
		slowPrev = slowPrev.Next
	}

	slowPrev.Next = slow.Next

	return dummy.Next
}
