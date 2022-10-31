package binary_tree_inorder_traversal

// 145. 二叉树的后序遍历
// https://leetcode.cn/problems/binary-tree-postorder-traversal/

// Node Definition for a binary tree node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func PostorderTraversal_Slice(root *Node) (res []int) {
	if root == nil {
		return
	}

	stack := []*Node{}
	cur := root
	prev := &Node{}

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if cur.Right == nil || cur.Right == prev {
			res = append(res, cur.Val)
			prev = cur
			cur = nil
		} else {
			stack = append(stack, cur)
			cur = cur.Right
		}
	}

	return
}
