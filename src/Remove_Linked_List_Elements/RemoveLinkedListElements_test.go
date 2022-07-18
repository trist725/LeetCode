package Remove_Linked_List_Elements

import "testing"

func BenchmarkRemoveElements_1(b *testing.B) {
	head := ListNode{}
	arr := []int{1, 2, 6, 3, 4, 5, 6}
	for i := 0; i < 7; i++ {
		tmp := ListNode{Val: arr[i]}
		head.Next = &tmp
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveElements_1(&head, 6)
	}
}

func BenchmarkRemoveElements_2(b *testing.B) {
	head := ListNode{}
	arr := []int{1, 2, 6, 3, 4, 5, 6}
	for i := 0; i < 7; i++ {
		tmp := ListNode{Val: arr[i]}
		head.Next = &tmp
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveElements_2(&head, 6)
	}
}
