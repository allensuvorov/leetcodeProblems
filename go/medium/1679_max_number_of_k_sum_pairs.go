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
    res := 0
    slices.Sort(nums)
    for l, r := 0, len(nums)-1; l<r; {
        if nums[l] + nums[r] < k {
            l++
        } else if nums[l] + nums[r] > k {
            r--
        } else {
            res++
            l++
            r--
        }
    }
    return res
}
