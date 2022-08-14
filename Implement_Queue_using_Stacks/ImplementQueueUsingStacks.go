package Implement_Queue_using_Stacks

// https://leetcode.cn/problems/implement-queue-using-stacks/

type MyQueue struct {
	In  []int
	Out []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.In = append(this.In, x)
}

func (this *MyQueue) in2out() {
	for len(this.In) > 0 {
		this.Out = append(this.Out, this.In[len(this.In)-1])
		this.In = this.In[:len(this.In)-1]
	}
}

func (this *MyQueue) Pop() int {
	if len(this.Out) == 0 {
		this.in2out()
	}
	x := this.Out[len(this.Out)-1]
	this.Out = this.Out[:len(this.Out)-1]
	return x

}

func (this *MyQueue) Peek() int {
	if len(this.Out) == 0 {
		this.in2out()
	}
	return this.Out[len(this.Out)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.In) == 0 && len(this.Out) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
