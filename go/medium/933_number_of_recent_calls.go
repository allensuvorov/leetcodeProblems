type RecentCounter struct {
    requests []int
    tail int
    head int
    len int
}


func Constructor() RecentCounter {
    len := 3000
    requests := make([]int, 10000)
    return RecentCounter{requests, 0, 0, len}
}


func (this *RecentCounter) Ping(t int) int {
    // Enqueue
    this.requests[this.tail] = t
    this.tail++

    // Dequeue requests that are older than 3000ms
    for this.requests[this.head] < t - this.len {
        this.head++
    }

    return this.tail - this.head
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
