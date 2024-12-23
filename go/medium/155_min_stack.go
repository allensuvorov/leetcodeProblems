type MinStack struct {
    stack []node
}

type node struct {
    val int
    minVal int
}

func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(val int)  {
    newNode := node{val:val, minVal:val}
    
    if len(this.stack) > 0 {
        newNode.minVal = min(val, this.stack[len(this.stack)-1].minVal)
    }
    this.stack = append(this.stack, newNode)
}


func (this *MinStack) Pop()  {
    this.stack = this.stack[:len(this.stack)-1]
}


func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1].val
}


func (this *MinStack) GetMin() int {
    return this.stack[len(this.stack)-1].minVal
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
