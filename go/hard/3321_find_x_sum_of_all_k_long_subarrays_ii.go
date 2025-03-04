/*
heaps and map is not enough
need a ordered set to access values directly
Go does not have ordred set, Btree is an option
might do later if appropriate
*/

import "container/heap"

func findXSum(nums []int, k int, x int) []int64 {
    // init two heaps and rebalance
    // top (to update xSum quickly) and bot
    // why use 

	h := &IntHeap{
		Nums:   make([]int, 0),
		Counts: map[int]int{},
	}

	for i := range k {
		if _, ok := h.Counts[nums[i]]; !ok {
			h.Nums = append(h.Nums, nums[i])
		}
		h.Counts[nums[i]]++
	}

	heap.Init(h)

	fmt.Println(h)
	n := len(nums)
	xSums := make([]int64, n-k+1)

	// get X-Sum
	// sliding window
	for i := range n - k + 1 {
		sum := 0
		for j := range x {
			num := h.Nums[j]
			sum += num * h.Counts[num]
		}
		xSums[i] = int64(sum)

		l := i
		r := i + k
		if r < len(nums) {
			h.Counts[nums[l]]--
			// binary search to find idx of [nums[l]]
			idx := binarySearch(h.Nums, h.Counts, nums[l])
            heap.Fix(h, idx)
			// fmt.Println(h)

			if _, ok := h.Counts[nums[r]]; !ok {
				heap.Push(h, nums[r])
				h.Counts[nums[r]]++
			} else {
				// binary search to find idx of [nums[r]]
				idx = binarySearch(h.Nums, h.Counts, nums[r])
				fmt.Println("fixing ", nums[r], "found at index", idx)
				h.Counts[nums[r]]++
				heap.Fix(h, idx)
			}
			fmt.Println(h)
		}

	}
	return xSums

}

func binarySearch(nums []int, counts map[int]int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		m := l + (r-l)/2
		guess := nums[m]
		if guess == target {
			return m
		}
		if counts[guess] > counts[target] {
			l = m + 1
		} else if counts[guess] < counts[target] {
			r = m - 1
		} else {
			if guess > target {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}
	return 0
}

// An IntHeap is a min-heap of ints.
type IntHeap struct {
	Nums   []int
	Counts map[int]int
}

func (h IntHeap) Len() int { return len(h.Nums) }

func (h IntHeap) Less(i, j int) bool {
	if h.Counts[h.Nums[i]] == h.Counts[h.Nums[j]] {
		return h.Nums[i] > h.Nums[j]
	}
	return h.Counts[h.Nums[i]] > h.Counts[h.Nums[j]]
}

func (h IntHeap) Swap(i, j int) { h.Nums[i], h.Nums[j] = h.Nums[j], h.Nums[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	h.Nums = append(h.Nums, x.(int))
}

func (h *IntHeap) Pop() any {
	old := h.Nums
	n := len(old)
	x := old[n-1]
	h.Nums = old[0 : n-1]
	return x
}

/*
import "container/heap"

func initializeHeaps(top *MinHeap, bot *MaxHeap, freqs map[int]int, x int) int {
	log.Println("heaps initialization start")
	sum := 0
	i := 0
	for value, priority := range freqs {
		(*bot)[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}

	heap.Init(bot)

	// move top x items the top heap and get initial xSum

	for len(*bot) > 0 && len(*top) < x {
		item := heap.Pop(bot).(*Item)
		sum += item.value * item.priority
		heap.Push(top, item)
	}

	heap.Init(top)

	log.Printf("heaps initialization - end len(top): %v, len(bot): %v \n", len(*top), len(*bot))
	return sum
}

func cutOldTail(top *MinHeap, bot *MaxHeap, freqs map[int]int, nums []int, oldTail, sum, x int) int {
	log.Println("cut oldTail - start")

	targetItem := &Item{
		value:    nums[oldTail],
		priority: freqs[nums[oldTail]],
	}

	// update heaps and freqs
	freqs[nums[oldTail]]--

	if freqs[nums[oldTail]] == 0 { // delete item
		log.Printf("cut oldTail delete value %v at index %v - start \n", nums[oldTail], oldTail)

		delete(freqs, nums[oldTail])

		if !less(targetItem, (*top)[0]) {
			log.Printf("cut oldTail delete value %v at index %v from top - start \n", nums[oldTail], oldTail)

			log.Printf("binarySearchTop, top - targetItem %+v \n", targetItem)

			index := binarySearchTop(*top, targetItem)
			sum -= (*top)[index].value
			heap.Remove(top, index)

			if len(*top) < x && len(*bot) > 0 {
				log.Printf("cut oldTail pop item from bot, len(bot): %v- start \n", len(*bot))
				item := heap.Pop(bot).(*Item)
				sum += item.value * item.priority
				heap.Push(top, item)
			}
		} else {
			index := binarySearchBot(*bot, targetItem)
			heap.Remove(bot, index)
		}
		log.Printf("cut oldTail delete value %v at index %v - end \n", nums[oldTail], oldTail)
	} else { // update item
		if !less(targetItem, (*top)[0]) { // from top
			index := binarySearchTop(*top, targetItem)
			(*top)[index].priority--
			sum -= (*top)[index].value
			heap.Fix(top, index)

			swap(top, bot, &sum)

		} else { // from bottom
			index := binarySearchBot(*bot, targetItem)
			(*bot)[index].priority = freqs[nums[oldTail]]
			heap.Fix(top, index)
		}
	}
	printTopHeap(*top)
	log.Println("cut oldTail - end")
	return sum
}

func addNewHead(top *MinHeap, bot *MaxHeap, freqs map[int]int, nums []int, newHead, sum, x int) int {
	log.Println("addNewHead - start")
	log.Printf("addNewHead %v at index %v", nums[newHead], newHead)
	log.Printf("addNewHead len(top) %v, len(bot) %v", len(*top), len(*bot))
	targetItem := &Item{
		value:    nums[newHead],
		priority: freqs[nums[newHead]],
	}

	freqs[nums[newHead]]++

	if freqs[nums[newHead]] == 1 { // add item
		targetItem.priority = 1

		if len(*top) < x || !less(targetItem, (*top)[0]) { // add to top
			heap.Push(top, targetItem)
			sum += targetItem.value

			if len(*top) > x {
				item := heap.Pop(top).(*Item) // move
				sum -= item.value * item.priority
				heap.Push(bot, item)
			}
		} else {
			heap.Push(bot, targetItem)
		}
	} else { // update item
		if !less(targetItem, (*top)[0]) { // in top

			index := binarySearchTop(*top, targetItem)
			(*top)[index].priority++
			sum += (*top)[index].value

            log.Println("\n")
			log.Println("heap.Fix - start - index:", index)
			
            printTopHeap(*top)
            
            heap.Fix(top, index)
            if !checkTopIsSortedIncreasing(*top) {
                fmt.Println("error - heap.Fix at line 126 - did not sort heap")
            }
            
            printTopHeap(*top)
            
			log.Println("heap.Fix - end")
            log.Println("\n")

			swap(top, bot, &sum)
		} else { // in bottom
			index := binarySearchBot(*bot, targetItem)
			(*bot)[index].priority = freqs[nums[newHead]]
			heap.Fix(bot, index)
			swap(top, bot, &sum)
		}
	}

	printTopHeap(*top)
	log.Println("addNewHead - end")
	return sum
}

func swap(top *MinHeap, bot *MaxHeap, sum *int) {
	log.Printf("addNewHead - swap - len(top) %v, len(bot) %v", len(*top), len(*bot))
	if len(*bot) > 0 && less((*top)[0], (*bot)[0]) { // swap
		item := heap.Pop(top).(*Item)
		*sum -= item.value * item.priority
		heap.Push(bot, item)
		fmt.Printf("swap from %v from top \n", item.value)

		item = heap.Pop(bot).(*Item)
		*sum += item.value * item.priority
		heap.Push(top, item)
		fmt.Printf("swap from %v from bot \n", item.value)
	}
}

func printTopHeap(top MinHeap) {
	log.Println("printTop - start, len is", len(top))
	for i := range top {
		log.Printf("top has: %v, %v, %v,", top[i].index, top[i].priority, top[i].value)
	}
}

func checkTopIsSortedIncreasing(top MinHeap) bool {
    isSorted := true
    for i := 1; i < len(top); i++ {
        if less(top[i], top[i-1]) {
            isSorted = false
        }
    }
    return isSorted
}

func findXSum(nums []int, k int, x int) []int64 {
	// two heaps, minHeap and maxHeap, with rebalancing
	// binary search
	// sliding window
	log.Println("findXSum start for ", nums)
    fmt.Println("findXSum start for ", nums)


	freqs := make(map[int]int)

	// initial k items for bottom heap
	for i := range k {
		freqs[nums[i]]++
	}

	top := make(MinHeap, 0)
	bot := make(MaxHeap, len(freqs))

	n := len(nums)
	xSums := make([]int64, n-k+1)

	xSums[0] = int64(initializeHeaps(&top, &bot, freqs, x))

	// sliding window
	for i := range n - k {
		sum := int(xSums[i])

		oldTail := i
		newHead := i + k

		// cut oldTail
		sum = cutOldTail(&top, &bot, freqs, nums, oldTail, sum, x)

		// newHead
		sum = addNewHead(&top, &bot, freqs, nums, newHead, sum, x)

		fmt.Println("freqs", freqs)

		xSums[i+1] = int64(sum)
	}
	log.Printf("findXSum end for %v with result %v \n", nums, xSums)
	return xSums
}

func binarySearchTop(pq MinHeap, target *Item) int {
	l := 0
	r := len(pq) - 1
	for l <= r {
		m := l + (r-l)/2
		guess := pq[m]
		if guess.value == target.value {
			return m
		}
		if guess.priority < target.priority {
			l = m + 1
		} else if guess.priority > target.priority {
			r = m - 1
		} else {
			if guess.value < target.value {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}
	return -1
}

func binarySearchBot(pq MaxHeap, target *Item) int {
	l := 0
	r := len(pq) - 1
	for l <= r {
		m := l + (r-l)/2
		guess := pq[m]
		if guess.value == target.value {
			return m
		}
		if guess.priority > target.priority {
			l = m + 1
		} else if guess.priority < target.priority {
			r = m - 1
		} else {
			if guess.value > target.value {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}
	return -1
}

// An Item is something we manage in a priority queue.
type Item struct {
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
	priority int // The priority of the item in the queue.
	value    int // The value of the item; arbitrary.
}

func less(a, b *Item) bool {
	if a.priority == b.priority {
		return a.value < b.value
	}
	return a.priority < b.priority
}

// A MinHeap implements heap.Interface and holds Items.
type MinHeap []*Item

func (pq MinHeap) Len() int { return len(pq) }

func (pq MinHeap) Less(i, j int) bool {
	log.Println("MinHeap.Less() - start")

	// We want Pop to give us the lowest, not highest, priority so we use less than here.
	if pq[i].priority == pq[j].priority {
		log.Printf("MinHeap.Less() - end - equal priority: %v < %v, result %t", pq[i], pq[j], pq[i].value < pq[j].value)
		return pq[i].value < pq[j].value
	}
	log.Printf("MinHeap.Less() - end - diff priority: %v < %v, result %t", pq[i], pq[j], pq[i].priority < pq[j].priority)
	return pq[i].priority < pq[j].priority
}

func (pq MinHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *MinHeap) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *MinHeap) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *MinHeap) update(item *Item, value int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

// A MaxHeap implements heap.Interface and holds Items.
type MaxHeap []*Item

func (pq MaxHeap) Len() int { return len(pq) }

func (pq MaxHeap) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	if pq[i].priority == pq[j].priority {
		return pq[i].value > pq[j].value
	}
	return pq[i].priority > pq[j].priority
}

func (pq MaxHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *MaxHeap) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *MaxHeap) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *MaxHeap) update(item *Item, value int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
*/
