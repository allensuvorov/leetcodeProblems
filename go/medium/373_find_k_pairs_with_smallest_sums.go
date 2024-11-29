import (
	"container/heap"
	"fmt"
)

type Pair struct {
    sum int
    i, j int
}

type IntHeap []Pair

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].sum < h[j].sum }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(Pair))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
    m := len(nums1)
    n := len(nums2)

    ans := [][]int{}
    visited := map[[2]int]bool{[2]int{0,0}:true}
    
    h := &IntHeap{{nums1[0] + nums2[0],0,0}}
    heap.Init(h)
    visited[[2]int{0,0}] = true

    for k > 0 && h.Len() > 0 {
        pair := heap.Pop(h).(Pair)
        i, j := pair.i, pair.j
        ans = append(ans, []int{nums1[i], nums2[j]})

        if i + 1 < m && !visited[[2]int{i + 1, j}] {
            heap.Push(h, Pair{nums1[i+1]+nums2[j], i + 1, j})
            visited[[2]int{i+1, j}] = true
        }

        if j + 1 < n && !visited[[2]int{i, j + 1}] {
            heap.Push(h, Pair{nums1[i]+nums2[j+1], i, j + 1})
            visited[[2]int{i, j+1}] = true
        }
        k = k - 1
    }
    return ans
}
