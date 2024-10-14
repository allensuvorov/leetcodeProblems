// sort
// prefix sum
// binary search
// Time : O(n * log n)
    
func answerQueries(nums []int, queries []int) []int {   
    slices.Sort(nums)
    for i := 1; i < len(nums); i++ {
        nums[i] += nums[i-1]
    }

    res := make([]int, len(queries))
    for i, v := range queries {
        res[i] = binarySearch(nums, v)
    }
    return res
}

func binarySearch(prefSums []int, target int) int {
    l, r := 0, len(prefSums)
    for l < r {
        m := l + (r - l)/2
        if prefSums[m] > target {
            r = m
        } else {
            l = m + 1
        }
    }
    return l
}
