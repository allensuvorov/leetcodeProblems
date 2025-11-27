type HitCounter struct {
    q []int
}


func Constructor() HitCounter {
    return HitCounter{[]int{}}
}


func (this *HitCounter) Hit(timestamp int)  {
    this.q = append(this.q, timestamp)
}


func (this *HitCounter) GetHits(timestamp int) int {
    for len(this.q) > 0 && this.q[0] <= timestamp - 300 {
        this.q = this.q[1:]
    }
    return len(this.q)
}




/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param_2 := obj.GetHits(timestamp);
 */
