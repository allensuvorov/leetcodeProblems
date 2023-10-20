func canBeIncreasing(nums []int) bool {
    limit := 1
    // check each adjacent pair
    // count non inc pairs
    count := 0
    for i := 0; i < len(nums) - 1 ; i ++ {
        if nums[i] >= nums[i + 1] {
            count++
        }
    }

    if count == 0 {
        return true
    }
    if count > limit {
        fmt.Println("more than 1 non inc")
        return false
    }
    // if there is one - remove one element, check near two
    for i := 0; i < len(nums) - 1 ; i ++ {
        if nums[i] >= nums[i + 1] {
            if i == 0 || i == len(nums) - 2 {
                return true
            }
            if nums[i-1] < nums[i+1] || nums[i] < nums[i+2]{
                return true
            }
            return false
        }
    }
    return true
}
