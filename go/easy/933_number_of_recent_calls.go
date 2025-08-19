type RecentCounter struct {
    timeLog []int
}


func Constructor() RecentCounter {
    return RecentCounter{[]int{}}
}


func (this *RecentCounter) Ping(t int) int {
    // append new ping
    this.timeLog = append(this.timeLog, t)
    
    //deque outdated pings
    for this.timeLog[0] < (t - 3000) {
        this.timeLog = this.timeLog[1:] 
    }
    return len(this.timeLog)
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
