package binary_tree_preorder_traversal

import "testing"

var root = &Node{}

func InitRoot() {
	cur := root
	for i := 0; i < 1000000; i++ {
		cur.Val = i
		cur.Left = &Node{}
		cur = cur.Left
	}
	root = &Node{}
	cur = root
	for i := 0; i < 1000000; i++ {
		cur.Val = i
		cur.Right = &Node{}
		cur = cur.Right
	}
}

func BenchmarkPreorderTraversal_List(b *testing.B) {
	InitRoot()
	b.ResetTimer()

	PreorderTraversal_List(root)
}

func BenchmarkPreorderTraversal_Slice(b *testing.B) {
	InitRoot()
	b.ResetTimer()

	PreorderTraversal_Slice(root)
}
