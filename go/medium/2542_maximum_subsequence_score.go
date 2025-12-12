import (
	"container/heap"
)

func maxScore(nums1 []int, nums2 []int, k int) int64 {
    n := len(nums1)
    pairs := make([][]int, n)
    for i := range pairs {
        pairs[i] = []int{nums1[i], nums2[i]}
    }

    slices.SortFunc(pairs, func(a, b []int) int { return b[1] - a[1]} )

    h := &minHeap{}
    heap.Init(h)
    var sum, res int
    for _, v := range pairs {
        heap.Push(h, v[0])
        sum += v[0]
        if h.Len() > k {
            sum -= heap.Pop(h).(int)
        }

        if h.Len() == k {
            res = max(res, sum * v[1])
        }
    }
    return int64(res)
}

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
