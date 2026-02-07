/*
- non pop vialations, so no mono Stack or mono Queue

input: ababc *** ab ****
output: bc
*/

func clearStars(s string) string {
    b := make(map[int]bool) // chars to be removed from the original
    var pq PriorityQueue
    heap.Init(&pq)
    for i := range s {
        if s[i] != '*' {
            heap.Push(&pq, &Item{s[i], i})
        } else {
            item := heap.Pop(&pq).(*Item)
            b[item.index] = true
        }
    }
    var res []byte
    for i := range s {
        if !b[i] && s[i] != '*' {
            res = append(res, s[i])
        }
    }
    return string(res)
}

type Item struct {
	value byte
	index int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
    if pq[i].value == pq[j].value {
        return pq[i].index > pq[j].index
    }
    return pq[i].value < pq[j].value
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
