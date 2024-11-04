import "container/heap"

func findXSum(nums []int, k int, x int) []int64 {
    
    // two heaps, minHeap and maxHeap, with rebalancing
    // binary search
    // sliding window

	top := make(MinHeap, x)
    bot := make(MaxHeap, k)
    
    freqs:= make(map[int]int)

    // initial k items for bottom heap
	for i := range k {
        freqs[nums[i]]++
	}

    i := 0
    for value, priority := range freqs{
        bot[i] = &Item{
            value: value,
            priority: priority,
            index: i,
        }
        i++
    }
    
    log.Println("heaps initialization start")

	heap.Init(&top)
    heap.Init(&bot)

    log.Println("heaps initialization end")

    // move top x items the top heap and get initial xSum
	n := len(nums)
	xSums := make([]int64, n-k+1)

    for range x {
        item := heap.Pop(&bot).(*Item)
        xSums[0] += int64(item.value)
        heap.Push(&top, item)
    }

	// sliding window
	for i := range n - k {
		sum := int(xSums[i])

		oldTail := i
        newHead := i + k
        
        targetItem := &Item{
            value: nums[oldTail],
            priority: freqs[nums[oldTail]],
        }
		
        // update heaps and freqs
        freqs[nums[oldTail]]--

        // cut oldTail
        if freqs[nums[oldTail]] == 0 { // delete item
            delete(freqs,nums[oldTail])
            if !less(targetItem, top[0]) {
                index := binarySearchTop(top, targetItem)
                sum -= top[index].value
                heap.Remove(&top, index)

                item := heap.Pop(&bot).(*Item)
                sum += item.value * item.priority
                heap.Push(&top, item)
            } else {
                index := binarySearchBot(bot, targetItem)
                heap.Remove(&bot, index)
            }
        } else { // update item
            if !less(targetItem, top[0]) { // from top
                index := binarySearchTop(top, targetItem)
                top[index].priority = freqs[nums[oldTail]]
                sum -= top[index].value
                heap.Fix(&top, i)
                if less(top[0], bot[0]) { // swap
                    item := heap.Pop(&top).(*Item)
                    sum -= item.value * item.priority
                    heap.Push(&bot, item)

                    item = heap.Pop(&bot).(*Item)
                    sum += item.value * item.priority
                    heap.Push(&top, item)
                }
            } else { // from bottom
                index := binarySearchBot(bot, targetItem)
                bot[index].priority = freqs[nums[oldTail]]
                heap.Fix(&top, index)
            }
        }

        // newHead
        targetItem = &Item{
            value: nums[newHead],
            priority: freqs[nums[newHead]],
        }

        freqs[nums[newHead]]++

        if freqs[nums[newHead]] == 1 { // add item
            targetItem.priority = 1

            if !less(targetItem, top[0]) { // add to top
                heap.Push(&top, targetItem)
                sum += targetItem.value

                item := heap.Pop(&top).(*Item) // move
                sum -= item.value * item.priority
                heap.Push(&bot, item)
            } else {
                heap.Push(&bot, targetItem)
            }
        } else { // update item
            if !less(targetItem, top[0]) { // from top
                index := binarySearchTop(top, targetItem)
                top[index].priority = freqs[nums[newHead]]
                sum += top[index].value
                heap.Fix(&top, i)
                if less(top[0], bot[0]) { // swap
                    item := heap.Pop(&top).(*Item)
                    sum -= item.value * item.priority
                    heap.Push(&bot, item)

                    item = heap.Pop(&bot).(*Item)
                    sum += item.value * item.priority
                    heap.Push(&top, item)
                }
            } else { // from bottom
                index := binarySearchBot(bot, targetItem)
                bot[index].priority = freqs[nums[oldTail]]
                heap.Fix(&top, index)
            }
        }

        xSums[i+1] = int64(sum)
	}
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
	value    int // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

func less(a, b *Item) bool {
	if a.priority == b.priority {
        return a.value < b.value
    }
    return a.priority < a.priority
}



// A MinHeap implements heap.Interface and holds Items.
type MinHeap []*Item

func (pq MinHeap) Len() int { return len(pq) }

func (pq MinHeap) Less(i, j int) bool {
    log.Println("MinHeap.Less() start")

	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	if pq[i].priority == pq[j].priority {
        return pq[i].value < pq[j].value
    }
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
