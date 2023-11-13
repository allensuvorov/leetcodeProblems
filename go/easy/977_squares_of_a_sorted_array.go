func sortedSquares(nums []int) []int {
    size := len(nums)
    res := make([]int, 0, size)
    // find the middle with binary search

    // for i := 0; nums[i] < 0 and i < size; i++ {
    //     middle = i
    // }

    r := binarySearch(nums)
    fmt.Println("start is at:", r)
    l := 0
    if r > 0 {
        l = r - 1
    }
    v := 0
    for l > 0 || r < size {
        rSq := nums[r]*nums[r]
        lSq := nums[l]*nums[l]
        if rSq <= lSq {
            v = rSq
            if r < size {
                r++
            }
        } else {
            v = lSq
            if l > 0 {
                l--
            } else {
                
            }
        }
        fmt.Println("square value is:", v, " pointers are:", l, r)
        res = append(res, v)
    }
    return res
}

func binarySearch(nums []int) int {
    l, r := 0, len(nums)
    for l < r {
        m := l + (r - l)/2 
        if nums[m] < 0 {
            l = m + 1
        } else {
            r = m
        }
    }
    return l
}
