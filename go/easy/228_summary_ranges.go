func summaryRanges(nums []int) []string {
    ranges := []string{}
    rng := ""
    l, r := 0, 0
    for i := 0; i < len(nums); i++ {
        // range start
        if rng == "" {
            l = nums[i]
            rng = rng + strconv.Itoa(l)
        }
        // range end
        if i == len(nums) - 1 || nums[i] + 1 != nums[i+1] {
            r = nums[i]
            if l != r {
                rng = strconv.Itoa(l) + "->" + strconv.Itoa(r)
            }
            ranges = append(ranges, rng)
            rng = ""
        }
    }
    return ranges
}
