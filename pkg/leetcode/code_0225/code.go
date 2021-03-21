package code_0225


type MyStack struct {
	queue1 []int // 正式队列
	queue2 []int // 辅助队列
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		queue1: make([]int, 0),
		queue2: make([]int, 0),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.queue2 = append(this.queue2, x)
	this.queue2 = append(this.queue2, this.queue1...)
	this.queue2, this.queue1 = this.queue1, this.queue2
	this.queue2 = make([]int, 0)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	// 题目假设 queue1 有元素
	result := this.queue1[0]
	this.queue1 = this.queue1[1:]
	return result
}

/** Get the top element. */
func (this *MyStack) Top() int {
	// 题目假设 queue1 有元素
	return this.queue1[0]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue1) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
