package Remove_Linked_List_Elements

import "testing"

func makeList() *ListNode {
	head := &ListNode{}
	arr := []int{0, 6, 1, 2, 3, 4, 5, 6}
	cur := &ListNode{Val: 666}
	head.Next = cur
	for i := 0; i < len(arr); i++ {
		tmp := ListNode{Val: arr[i]}
		cur.Next = &tmp
		cur = cur.Next
	}
	return head
}

func BenchmarkRemoveElements_1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveElements_1(makeList(), 6)
	}
}

func BenchmarkRemoveElements_2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveElements_2(makeList(), 6)
	}
}
