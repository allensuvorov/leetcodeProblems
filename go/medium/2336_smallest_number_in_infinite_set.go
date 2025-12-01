import (
	"container/heap"
)

type SmallestInfiniteSet struct {
    isPresent map[int]bool
    h *IntHeap
    currentInteger int
}


func Constructor() SmallestInfiniteSet {
    h := &IntHeap{}
    heap.Init(h)
    return SmallestInfiniteSet{
        isPresent: make(map[int]bool),
        h: h,
        currentInteger: 1,
    }
}


func (this *SmallestInfiniteSet) PopSmallest() int {
    if this.h.Len() == 0 {
        old := this.currentInteger
        this.currentInteger++
        return old
    }
    delete(this.isPresent, (*this.h)[0])
    return heap.Pop(this.h).(int)
}


func (this *SmallestInfiniteSet) AddBack(num int)  {
    if num < this.currentInteger && !this.isPresent[num] {
        this.isPresent[num] = true
        heap.Push(this.h, num)
    }
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */





