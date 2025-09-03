type MyQueue struct {
	revSt []int // reversing stack
	popSt []int // popping stack
}

func Constructor() MyQueue {
	return MyQueue{[]int{}, []int{}}
}

func (this *MyQueue) Push(x int) {
	this.revSt = append(this.revSt, x)
}

func (this *MyQueue) Pop() int {
    res := this.Peek()

	if len(this.popSt) != 0 {
		this.popSt = this.popSt[:len(this.popSt)-1]
	}

	return res
}

func (this *MyQueue) Peek() int {
    if this.Empty() {
        return 0
    }

	if len(this.popSt) == 0 { // only if the pop stack is empty
    	// pop from revSt, push to popSt
		for len(this.revSt) > 0 {
            topVal := this.revSt[len(this.revSt)-1]
            this.revSt = this.revSt[:len(this.revSt)-1]
            this.popSt = append(this.popSt, topVal)
        }
	}

	return this.popSt[len(this.popSt)-1]
}

func (this *MyQueue) Empty() bool {
    return len(this.revSt) == 0 && len(this.popSt) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
