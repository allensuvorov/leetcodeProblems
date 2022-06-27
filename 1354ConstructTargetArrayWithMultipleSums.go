// 1354. Construct Target Array With Multiple Sums
// ideas: recursion, will it help with time complexity?
// ideas: heap
package main

import (
	"fmt"
)
// heapify slice into maxHeap
func maxHeapify(heap *[]int, i int) {
  max := i
  lChild := 2*i + 1
  rChild := 2*i + 2

  if lChild < len(*heap) && (*heap)[lChild] > (*heap)[max] {
    max = lChild
  }
  if rChild < len(*heap) && (*heap)[rChild] > (*heap)[max] {
    max = rChild
  }

  if max != i {
    (*heap)[i], (*heap)[max] = (*heap)[max], (*heap)[i]
    maxHeapify(heap, max)
  }
}

func buildHeap(heap []int) {
  n := len(heap)
  startIdx := n / 2 - 1

  for i := startIdx; i >= 0; i-- {
    maxHeapify(&heap, i)
  }
}

func isPossible(target []int) bool {
	max := 0
	sum := 0  
	maxI := 0
	rest := 0 // rest of array, excluding max
	
  if len(target) == 1 {
		if target[0] == 1 {
			return true
		} else {
			return false
		}
	}

	for _, num := range target {
		sum += num
	}

  buildHeap(target)
  fmt.Println(target)
  
	for {
    max = target[0]
    rest = sum - max

		if max == 1 || rest == 1 {
			return true
		}

		if max < rest+1 || max%rest == 0 {
			return false
		} else {
			target[maxI] = max % rest // replace max with remainder
		}
		sum -= max - target[maxI] // update sum
    maxHeapify(&target, 0)
	}
}

func main() {

	fmt.Println(isPossible([]int{1, 1}))
	fmt.Println(isPossible([]int{2, 900000001}))
	fmt.Println(isPossible([]int{2}))
  	fmt.Println(isPossible([]int{1,2,3,4,5,6,7,8}))
  	fmt.Println(isPossible([]int{1,1,1,7}))
}
