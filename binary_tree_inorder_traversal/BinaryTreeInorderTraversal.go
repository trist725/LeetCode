package binary_tree_inorder_traversal

// 94. 二叉树的中序遍历
// https://leetcode.cn/problems/binary-tree-inorder-traversal/

// Node Definition for a binary tree node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func InorderTraversal_Slice(root *Node) (res []int) {
	if root == nil {
		return
	}

	stack := []*Node{}
	cur := root

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		res = append(res, stack[len(stack)-1].Val)
		cur = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}

	return
}
