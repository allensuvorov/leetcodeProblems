func findKthLargest(nums []int, k int) int {
    h := &maxHeap{}
    heap.Init(h)
    for i := range nums {
        heap.Push(h, nums[i])
    }

    res := 0
    for range k {
        res = heap.Pop(h).(int)
    }
    return res
}

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
