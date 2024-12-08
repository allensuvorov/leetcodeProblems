type MedianFinder struct {
    minH *minHeap
    maxH *maxHeap
}


func Constructor() MedianFinder {
    return MedianFinder {&minHeap{}, &maxHeap{}}
}


func (this *MedianFinder) AddNum(num int)  {
    if len(*this.minH) > 0 && num >= (*this.minH)[0] { // (*this.maxH)[0]
        heap.Push(this.minH, num)
    } else {
        heap.Push(this.maxH, num)
    }
    this.rebalance()
}


func (this *MedianFinder) FindMedian() float64 {
    if len(*this.minH) == len(*this.maxH) {
        return float64(float64((*this.minH)[0] + (*this.maxH)[0])/2)
    } else {
        return float64((*this.maxH)[0])
    }
}


func (this *MedianFinder) rebalance() { // len of max has to be equal of longer by one
    if len(*this.minH) - len(*this.maxH) == 1 { // min is longer by one
        num := heap.Pop(this.minH).(int)
        heap.Push(this.maxH, num)
    } else if len(*this.maxH) - len(*this.minH) > 1 { // max is longer by more than one
        num := heap.Pop(this.maxH).(int)
        heap.Push(this.minH, num)
    }
}

// An maxHeap is a max-heap of ints.
type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// An minHeap is a min-heap of ints.
type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
