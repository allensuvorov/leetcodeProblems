type SmallestInfiniteSet struct {
    h *minHeap
    set map[int]bool
}


func Constructor() SmallestInfiniteSet {
    h := &minHeap{}
    heap.Init(h)
    set := make(map[int]bool)
    for i := 1; i <= 1000; i++ {
        heap.Push(h, i)
        set[i] = true
    }
    return SmallestInfiniteSet{h, set}
}


func (this *SmallestInfiniteSet) PopSmallest() int {
    num := heap.Pop(this.h).(int)
    this.set[num] = false
    return num
}


func (this *SmallestInfiniteSet) AddBack(num int)  {
    if !this.set[num] {
        heap.Push(this.h, num )
        this.set[num] = true
    }
}


/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */

 type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
