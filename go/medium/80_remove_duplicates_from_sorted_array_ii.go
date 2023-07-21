func removeDuplicates(nums []int) int {
    limit, countDups := 1, 0
    // 2 pointers: collector + scout
    collector, scout := 0, 1
    collect := func(){
        collector++
        nums[collector] = nums[scout]
    }
    for ; scout < len(nums); scout++ {
        if nums[scout] == nums[collector] {
            if countDups < limit {
                collect()
            }
            countDups++
        } else {
            collect()
            countDups = 0
        }
    }
    return collector + 1
}
