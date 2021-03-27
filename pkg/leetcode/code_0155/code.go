package code_0155

// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
// 常数时间内检索

type MinStack struct {
	stackNodeList []*stackNode
}

type stackNode struct {
	minValue int
	value    int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{stackNodeList: make([]*stackNode, 0)}
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (this *MinStack) Push(val int) {
	minValue := 0
	if len(this.stackNodeList) == 0 {
		minValue = val
	} else {
		minValue = min(this.GetMin(), val)
	}

	this.stackNodeList = append(this.stackNodeList, &stackNode{
		minValue: minValue,
		value:    val,
	})
}

func (this *MinStack) Pop() {
	if len(this.stackNodeList) == 0 {
		return
	}
	this.stackNodeList = this.stackNodeList[:len(this.stackNodeList)-1]
	return
}

func (this *MinStack) Top() int {
	if len(this.stackNodeList) == 0 {
		return 0
	}
	return this.stackNodeList[len(this.stackNodeList)-1].value
}

func (this *MinStack) GetMin() int {
	if len(this.stackNodeList) == 0 {
		return 0
	}
	return this.stackNodeList[len(this.stackNodeList)-1].minValue
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
