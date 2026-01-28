type RLEIterator struct {
    i int // 
    j int // 
    encoding []int
}


func Constructor(encoding []int) RLEIterator {
    return RLEIterator{0, 0, encoding}
}


func (this *RLEIterator) Next(n int) int {
    repeated := this.encoding[i]
    number := this.encoding[i+1]
    
    this.j++ // move to next

    if this.j > repeated {
        if this.i == len(this.encoding) - 2 {
            return -1
        }
            this.i++
        }
    }



}


/**
 * Your RLEIterator object will be instantiated and called as such:
 * obj := Constructor(encoding);
 * param_1 := obj.Next(n);
 */

 /*
 
j = 0
repeated := 0

 */
