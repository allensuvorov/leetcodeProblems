import "container/heap"

func maxScore(nums1 []int, nums2 []int, k int) int64 {
    pairs := make([][]int, len(nums1))
    for i := range pairs {
        pairs[i] = []int{nums1[i], nums2[i]}
    }

    slices.SortFunc(pairs, func(a, b []int)int {
        return b[1] - a[1]
    })
    
    ans := 0
    sum := 0
    h := &IntHeap{}
    heap.Init(h)

    for _, v := range pairs {
        
        sum += v[0]
        heap.Push(h, v[0])

        if len(*h) > k {
            sum -= heap.Pop(h).(int)
        }

        if len(*h) == k {
            ans = max(ans, sum * v[1])
        }
    }
    return int64(ans)
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
