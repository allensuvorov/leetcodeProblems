type MyStack struct {
    q []int // <- queue direction
}


func Constructor() MyStack {
    return MyStack{[]int{}}
}

// reversal is possible 
// cycle the queue n-1 times
// easy with two entries: ab -> ba
// with three just abstract: q = "ab"
// qc - > cq with q cycles

func (this *MyStack) Push(x int)  { // O(n) 
    this.q = append(this.q, x)
    for range len(this.q) - 1 {
        front := this.Pop()
        this.q = append(this.q, front)
    }
}


func (this *MyStack) Pop() int { // O(1)
    result := this.q[0]
    this.q = this.q[1:]
    return result
}


func (this *MyStack) Top() int {
    return this.q[0]
}


func (this *MyStack) Empty() bool {
    return len(this.q) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
