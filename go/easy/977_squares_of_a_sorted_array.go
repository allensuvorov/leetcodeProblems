func sortedSquares(nums []int) []int {
    size := len(nums)
    res := make([]int, 0, size)

    r := binarySearch(nums)
    //fmt.Println("start is at:", r)
    
    l := -1
    if r > 0 {
        l = r - 1
    }

    v := 0
    for l >= 0 && r < size {
        rSq := nums[r]*nums[r]
        lSq := nums[l]*nums[l]
        if rSq <= lSq {
            v = rSq
            if r < size {
                r++
            }
        } else {
            v = lSq
            if l >= 0 {
                l--
            }
        }
        //fmt.Println("square value is:", v, " pointers are:", l, r)
        res = append(res, v)
    }

    for r < size {
        v = nums[r]*nums[r]
        res = append(res, v)
        r++
    }
    for l >= 0 {
        v = nums[l]*nums[l]
        res = append(res, v)
        l--
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
