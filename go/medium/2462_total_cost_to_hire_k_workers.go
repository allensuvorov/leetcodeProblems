import (
	"container/heap"
)

func totalCost(costs []int, k int, candidates int) int64 {
    lHeap := &IntHeap{}
    rHeap := &IntHeap{}
    
    l := candidates - 1
    r := max(candidates, len(costs) - candidates)
    
    *lHeap = costs[:l + 1]
    *rHeap = costs[r:]

    heap.Init(lHeap)
    heap.Init(rHeap)
    
    res := 0

    for range k {
        if (rHeap.Len() != 0 && lHeap.Len() != 0 && (*rHeap)[0] < (*lHeap)[0]) || 
            lHeap.Len() == 0 {
            res += heap.Pop(rHeap).(int)
            if r - l > 1 {
                r--
                heap.Push(rHeap, costs[r])
            }
        } else {
            res += heap.Pop(lHeap).(int)
            if r - l > 1 {
                l++
                heap.Push(lHeap, costs[l])
            }
        }
    }
    return int64(res)
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
