func rob(nums []int) int {

    curMax := 0
    prevMax := 0

    for _, stash := range nums {
        option1 := prevMax + stash
        option2 := curMax

        nextMax := max(option1, option2)

        prevMax = curMax
        curMax = nextMax
    }
    return curMax
}

// 1, 2, 3
// prev, cur, next
