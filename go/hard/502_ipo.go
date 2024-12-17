type project struct {
    profit int
    capital int
}

type PriorityQueue []*project

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].profit > pq[j].profit }
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x any) {
	item := x.(*project)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	*pq = old[0 : n-1]
	return item
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
    if len(profits) == 0 || len(capital) == 0 {
        return 0
    }
    projects := make([]*project, len(profits))
    
    for i := range projects {
        projects[i] = &project{profits[i], capital[i]}
    }

    slices.SortFunc(projects, func(a, b *project) int {
        return a.capital - b.capital
    })

    pq := make(PriorityQueue, 0)
    heap.Init(&pq)
    
    i := 0
    for range k {
        for i < len(projects) && projects[i].capital <= w { // enque if we can afford it
            heap.Push(&pq, projects[i])
            i++
        }

        if len(pq) == 0 {
            break
        }

        maxProfitProject := heap.Pop(&pq).(*project)
        w += maxProfitProject.profit
    }

    return w
}
