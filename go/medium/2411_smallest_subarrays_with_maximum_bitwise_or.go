func smallestSubarrays(nums []int) []int {
    maxOr := 0
    for _, v := range nums {
        maxOr |= v
    }   

    // sliding window
    res := make([]int, len(nums))
    curOr := 0
    for l, r := 0, 0; r < len(nums); r++ {
        curOr |= nums[r]
        if curOr == maxOr {
            res[l] = r - l + 1
            
            l++
        }
    } 
}
