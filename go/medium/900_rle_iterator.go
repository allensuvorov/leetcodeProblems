type RLEIterator struct {
    i int
    a []int
}


func Constructor(a []int) RLEIterator {
    return RLEIterator{0, a}
}


func (this *RLEIterator) Next(n int) int {
    for this.i < len(this.a) {
        if n > this.a[this.i] {
            n -= this.a[this.i]
            this.i += 2
        } else {
            this.a[this.i] -= n
            return this.a[this.i+1]
        }
    }
    return -1
}


/**
 * Your RLEIterator object will be instantiated and called as such:
 * obj := Constructor(a);
 * param_1 := obj.Next(n);
 */

 /*
 2,a

 */
