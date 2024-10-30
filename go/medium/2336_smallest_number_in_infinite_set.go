import "container/heap"

type SmallestInfiniteSet struct {
    H *IntHeap
    inHeap map[int]bool
}


func Constructor() SmallestInfiniteSet {
    h := make(IntHeap, 1000)
    inHeap := map[int]bool{}
    for i := range h {
        h[i] = i + 1
        inHeap[i] = true
    }
	heap.Init(&h)
    return SmallestInfiniteSet{&h,inHeap}
}


func (this *SmallestInfiniteSet) PopSmallest() int {
    num := heap.Pop(this.H).(int)
    this.inHeap[num] = false
    return num
}


func (this *SmallestInfiniteSet) AddBack(num int)  {
    if !this.inHeap[num] {
        this.inHeap[num] = true
        heap.Push(this.H, num)
    }
}


/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */

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
