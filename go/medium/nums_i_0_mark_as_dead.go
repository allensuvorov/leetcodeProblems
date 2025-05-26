func circularArrayLoop(nums []int) bool {
    for i, v := range nums {
        if v % len(nums) != 0 { // if not a dead end
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
    if nums[i] > 1000 { // if visited
        return true
    }
    
    if nums[i] % len(nums) == 0 { // if it is a dead end
        return false
    }
    
    if nums[i] * rootSign < 0 { // if nums[i] sign opposit then root sign
        return false
    }
    
    nums[i] += 5000 // mark as visited

    if dfs(nums, (i + nums[i])%len(nums), rootSign) { // recursive call
        return true
    } 
    nums[i] = 0 // mark as dead
    return false
}

// [2,-1,1,2,2]
//  ^
//  rootSign = 1
// call - nums[0] = 2: return true


// check each node to look for a loop
// disqual 1 - for a opposite sign (mark as deadend all but the last)
// disqual 2 - a one node loop at the end (mark as deadend all but the last)
// visited 
// once a qualified cycle is confirmed - done

