package Remove_Linked_List_Elements

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// https://leetcode.cn/problems/remove-linked-list-elements/

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveElements_1(head *ListNode, val int) *ListNode {
	var (
		prev = head
		cur  = head
	)

	for cur != nil {
		if cur.Val == val {
			if cur == head {
				head = cur.Next
				prev = head
			} else {
				prev.Next = cur.Next
			}
		} else { // cur被删则不能移动prev
			prev = cur
		}
		cur = cur.Next
	}

	return head
}

// 虚拟头节点
func RemoveElements_2(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	cur := dummyHead
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}
