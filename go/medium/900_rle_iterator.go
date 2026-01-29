type RLEIterator struct {
    i int
    encoding []int
}


func Constructor(encoding []int) RLEIterator {
    return RLEIterator{0, encoding}
}


func (this *RLEIterator) Next(n int) int {
    res := -1
    for n > 0 {
        // zero hopper
        for this.i < len(this.encoding) && this.encoding[this.i] == 0 {
            this.i += 2
        }

        if this.i >= len(this.encoding) {
            return -1
        }

        res = this.encoding[this.i+1]
        
        drain := min(n, this.encoding[this.i])
        n -= drain
        this.encoding[this.i] -= drain
    }
    return res
}


/**
 * Your RLEIterator object will be instantiated and called as such:
 * obj := Constructor(encoding);
 * param_1 := obj.Next(n);
 */
