package Design_Linked_List

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

// 707. 设计链表
// https://leetcode.cn/problems/design-linked-list/

// 笑了 这里struct名字用ListNode力扣会编译报错:
// Undefined Deserializer solution.go
type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

type MyLinkedList struct {
	Dummy *Node
	Size  int
}

func Constructor() MyLinkedList {
	return MyLinkedList{Dummy: &Node{}}
}

// index从0开始
func (this *MyLinkedList) Get(index int) int {
	count := 0
	cur := this.Dummy
	for cur.Next != nil {
		cur = cur.Next
		if count == index {
			return cur.Val
		}
		count++
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	next := this.Dummy.Next

	this.Dummy.Next = &Node{Val: val, Prev: this.Dummy, Next: next}
	if next != nil {
		next.Prev = this.Dummy.Next
	}
	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	cur := this.Dummy
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &Node{Val: val, Prev: cur}
	this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == this.Size {
		this.AddAtTail(val)
	} else if index < 0 {
		this.AddAtHead(val)
	} else if index > this.Size {
		// do nothing
	} else { // >=0 && <size
		count := 0
		cur := this.Dummy
		for cur.Next != nil {
			cur = cur.Next
			if count == index {
				newNode := &Node{Val: val, Prev: cur.Prev, Next: cur}
				cur.Prev.Next = newNode
				cur.Prev = newNode
				this.Size++
				return
			}
			count++
		}
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index > this.Size {
		return
	}

	count := 0
	cur := this.Dummy
	for cur.Next != nil {
		cur = cur.Next
		if count == index {
			cur.Prev.Next = cur.Next
			if cur.Next != nil {
				cur.Next.Prev = cur.Prev
			}
			this.Size--
			return
		}
		count++
	}
}
