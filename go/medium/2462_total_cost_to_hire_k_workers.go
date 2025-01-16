func totalCost(costs []int, k int, candidates int) int64 {
    h := &minHeap{}
    heap.Init(h)
    l, r := 0, len(costs) - 1
    for i := 0; i < candidates && l <= r; i++ {
        if l <= r {
            heap.Push(h, candidate{costs[l], l})
            l++
        }
        if l <= r { 
            heap.Push(h, candidate{costs[r], r})
            r--
        }
    }

    var result int64
    for range k {
        cand := heap.Pop(h).(candidate)
        result += int64(cand.cost)
        if l <= r {
            if cand.index <= l {
                heap.Push(h, candidate{costs[l], l})
                l++
            } else {
                heap.Push(h, candidate{costs[r], r})
                r--
            }
        }
    }
    return result
}

type candidate struct {
    cost int
    index int
}

type minHeap []candidate

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { 
    if h[i].cost == h[j].cost {
        return h[i].index < h[j].index
    }
    return h[i].cost < h[j].cost
}

func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(candidate))
}

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
