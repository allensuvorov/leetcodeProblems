// hash map
func maxOperations(nums []int, k int) int {
    counts := map[int]int{}
    res := 0
    for _, v := range nums {
        if counts[k-v] > 0 {
            res++
            counts[k-v]--
        } else {
            counts[v]++
        }
    }
    return res
}


// sorting
func maxOperations(nums []int, k int) int {
    slices.Sort(nums)
    l := 0
    r := len(nums) - 1
    result := 0
    for l < r {
        switch {
        case nums[l] + nums[r] == k:
            result++
            l++
            r--
        case nums[l] + nums[r] > k:
            r--
        case nums[l] + nums[r] < k:
            l++
        }
    }
    return result
}
