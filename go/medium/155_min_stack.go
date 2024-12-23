type MinStack struct {
    vs []int
    ms []int
}


func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(val int)  {
    this.vs = append(this.vs, val)

    if len(this.ms) == 0 || val <= this.ms[len(this.ms) - 1] {
        this.ms = append(this.ms, val)
    }
}


func (this *MinStack) Pop()  {
    if this.vs[len(this.vs)-1] == this.ms[len(this.ms)-1] {
        this.ms = this.ms[:len(this.ms)-1]
    }
    this.vs = this.vs[:len(this.vs)-1]
}


func (this *MinStack) Top() int {
    return this.vs[len(this.vs)-1]
}


func (this *MinStack) GetMin() int {
    return this.ms[len(this.ms)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
