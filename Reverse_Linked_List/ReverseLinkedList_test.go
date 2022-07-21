package Reverse_Linked_List

import "testing"

func makeList() *ListNode {
	head := &ListNode{}
	arr := []int{1, 2, 3, 4, 5}
	for i := 0; i < 10; i++ {
		arr = append(arr, arr...)
	}
	cur := &ListNode{Val: 666}
	head.Next = cur
	for i := 0; i < len(arr); i++ {
		tmp := ListNode{Val: arr[i]}
		cur.Next = &tmp
		cur = cur.Next
	}
	return head
}

func BenchmarkReverseList_1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReverseList_1(makeList())
	}
}

func BenchmarkReverseList_2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReverseList_2(makeList())
	}
}
