package binary_tree_preorder_traversal

import "container/list"

// 144. 二叉树的前序遍历
// https://leetcode.cn/problems/binary-tree-preorder-traversal/

// Node Definition for a binary tree node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func PreorderTraversal_Slice(root *Node) (res []int) {
	if root == nil {
		return
	}

	l := []*Node{}
	cur := root

	for cur != nil || len(l) > 0 {
		for cur != nil {
			res = append(res, cur.Val)
			l = append(l, cur)
			cur = cur.Left
		}

		cur = l[len(l)-1].Right
		l = l[:len(l)-1]
	}

	return
}

func PreorderTraversal_List(root *Node) (res []int) {
	if root == nil {
		return
	}

	l := list.New()
	cur := root

	for cur != nil || l.Len() > 0 {
		for cur != nil {
			res = append(res, cur.Val)
			l.PushBack(cur)
			cur = cur.Left
		}

		cur = l.Back().Value.(*Node).Right
		l.Remove(l.Back())
	}

	return
}
