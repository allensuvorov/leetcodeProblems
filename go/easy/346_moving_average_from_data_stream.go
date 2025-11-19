type MovingAverage struct {
    size int
    sum int
    q []int
}


func Constructor(size int) MovingAverage {
    return MovingAverage{
        size: size,
        sum: 0,
        q: make([]int, 0, size),
    }
}


func (this *MovingAverage) Next(val int) float64 {
    if len(this.q) == this.size {
        this.sum -= this.q[0]
        this.q = this.q[1:]
    }
    this.sum += val
    this.q = append(this.q, val)

    return float64(this.sum) / float64(len(this.q))
}


/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
