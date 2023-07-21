func removeDuplicates(nums []int) int {
    // 2 pointers: collector + scout
    collector := 0
    limit := 1
    countDup := 0
    for scout := 1; scout < len(nums); scout++ {
        if nums[scout] == nums[collector] {
            if countDup < limit {
                nums[collector + 1] = nums[scout]
                collector++
            }
            countDup++
        } else {
            nums[collector + 1] = nums[scout]
            collector++
            countDup = 0
        }
    }
    return collector + 1
}
