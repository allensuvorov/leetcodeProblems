type RecentCounter struct {
    q []int    
}


func Constructor() RecentCounter {
    return RecentCounter{
        q: []int{},
        }
}


func (this *RecentCounter) Ping(t int) int {
    this.q = append(this.q, t)
    for t - 3000 > this.q[0] { // 3002 - 3000 > 1
        this.q = this.q[1:]
    }
    return len(this.q)
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
