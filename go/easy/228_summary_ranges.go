func summaryRanges(nums []int) []string {
    l := 0
    inRange := false
    res := []string{}
    for i := range nums {
        if inRange {
            if i == len(nums) - 1 || nums[i+1] - nums[i] > 1 {
                res = append(res, strconv.Itoa(l)+"->"+strconv.Itoa(nums[i]))
                l = 0
                inRange = false
            }
        } else {
            if i == len(nums) - 1 || nums[i+1] - nums[i] > 1 {
                res = append(res, strconv.Itoa(nums[i]))
            } else {
                l = nums[i]
                inRange = true
            }
        }  
    }
    return res
}
