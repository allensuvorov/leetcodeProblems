func totalCost(costs []int, k int, candidates int) int64 {
    var ans int64 = 0
    n := len(costs)
    l, r := 0, n - 1

    // initialize pq
    pq := PriorityQueue{}
    for l < min(r, candidates) {
        pq = append(pq, &Cost{costs[l], l})
        l++
    }
    l--
    for r > max(l, n - candidates - 1) {
        pq = append(pq, &Cost{costs[r], r})
        r--
    }
    r++
    heap.Init(&pq)

    for range k {
        // hire a worker
        hired := heap.Pop(&pq).(*Cost)
        // fmt.Printf("Hired worker: %+v, l=%v, r=%v, n=%v \n", hired, l, r, n)
        ans += int64(hired.value)
        if l + 1 < r {
            if hired.index <= l {
                l++
                heap.Push(&pq, &Cost{costs[l], l})
            } else {
                r--
                heap.Push(&pq, &Cost{costs[r], r})
            }
        }
    }

    return ans
}


type Cost struct {
	value    int
	index   int
}

type PriorityQueue []*Cost

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
    if pq[i].value == pq[j].value {
        return pq[i].index < pq[j].index
    }
	return pq[i].value < pq[j].value
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	cost := x.(*Cost)
	*pq = append(*pq, cost)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	cost := old[n-1]
	*pq = old[0 : n-1]
	return cost
}
