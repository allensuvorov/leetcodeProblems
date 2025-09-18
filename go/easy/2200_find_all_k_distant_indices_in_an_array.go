func findKDistantIndices(nums []int, key int, k int) []int {
    // 0 ... n - 1
    n := len(nums)

    keyIndexList := []int{}
    for i := range nums {
        if nums[i] == key {
            keyIndexList = append(keyIndexList, i)
        }
    }
    
    result := []int{}

    for _, i := range keyIndexList {
        l := i - k
        l = max(0, l)
        if len(result) > 0 {
            topVal := result[len(result) - 1]
            l = max(l, topVal + 1)
        }

        r := i + k
        r = min(n-1, r)

        for j := l; j <= r; j++ {
            result = append(result, j)
        }
    }
    return result
}
