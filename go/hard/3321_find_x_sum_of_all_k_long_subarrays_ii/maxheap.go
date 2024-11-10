package main

// // A MaxHeap implements heap.Interface and holds Items.
// type MaxHeap []*Item

// func (pq MaxHeap) Len() int { return len(pq) }

// func (pq MaxHeap) Less(i, j int) bool {
// 	if pq[i].priority == pq[j].priority {
// 		return pq[i].value > pq[j].value
// 	}
// 	return pq[i].priority > pq[j].priority
// }

// func (pq MaxHeap) Swap(i, j int) {
// 	pq[i], pq[j] = pq[j], pq[i]
// }

// func (pq *MaxHeap) Push(x any) {
// 	*pq = append(*pq, x.(*Item))
// }

// func (pq *MaxHeap) Pop() any {
// 	old := *pq
// 	n := len(old)
// 	item := old[n-1]
// 	old[n-1] = nil // don't stop the GC from reclaiming the item eventually
// 	*pq = old[0 : n-1]
// 	return item
// }
