func circularArrayLoop(nums []int) bool {
    for i, v := range nums {
        // if not a dead end
        if v % len(nums) != 0 { 
            rootSign := 1
            if v < 0 {
                rootSign = -1
            }
            if dfs(nums, i, rootSign) {
                return true
            }
        }
    }
    return false
}

func dfs(nums []int, i int, rootSign int) bool {
    // if visited
    if nums[i] > 1000 { 
        return true
    }
    // if dead end OR nums[i] sign opposit then root sign
    if nums[i] % len(nums) == 0 || nums[i] * rootSign < 0 { 
        return false
    }
    // next is 
    next := (i + nums[i])%len(nums)
    if next < 0 {
        next = len(nums) + next // revolve back from end
    }
    // mark as visited
    nums[i] += 5000 
    // recursive call
    if dfs(nums, next, rootSign) { 
        return true
    } 
    // mark as deadend
    nums[i] = 0 
    return false
}
