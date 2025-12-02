func longestSubarray(nums []int, limit int) int {
    ans := 0
    minQue := list.New()
    maxQue := list.New()
    l := 0
    for r, v := range nums {
        // drain the backs
        for minQue.Len() > 0 {
            back := minQue.Back().Value.(node)
            if nums[r] < back.value {
                minQue.Remove(minQue.Back())
            } else {
                break
            }
        }
        for maxQue.Len() > 0 {
            back := maxQue.Back().Value.(node)
            if nums[r] > back.value {
                maxQue.Remove(maxQue.Back())
            } else {
                break
            }
        }
        minQue.PushBack(node{r, v})
        maxQue.PushBack(node{r, v})
        for {
            drainFront(minQue, l)
            drainFront(maxQue, l)
            if maxQue.Front().Value.(node).value - minQue.Front().Value.(node).value > limit {
                l++
            } else {
                break
            }
        }
        ans = max(ans, r - l + 1)
    }
    return ans
}

func drainFront (q *list.List, l int) {
    for q.Len() > 0 {
        front := q.Front().Value.(node)
        if front.index < l {
            q.Remove(q.Front())
        } else {
            break
        }

    }
}

type node struct {
    index int
    value int
}
