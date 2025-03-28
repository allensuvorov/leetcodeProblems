func rearrangeArray(nums []int) []int {
    n := len(nums)
    res := make([]int, 0, len(nums))
    
    neg, pos := 0, 0; 

    for neg < n && pos < n {
        // search for next positive
        for pos < n && nums[pos] < 0 {
            pos++
        }
        res = append(res, nums[pos])
        pos++
        
        // search for next negative
        for neg < n && nums[neg] > 0 {
            neg++
        }
        res = append(res, nums[neg])
        neg++
    }
    return res
}
