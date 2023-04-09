func findMaxK(nums []int) int {

// every number placed it the index+1
// collect nums inside the array
// shift -1 to index

    for i := range nums {
        // head
        av := abs(nums[i])
        if av - 1 == i || av - 1 == 0 {
            continue
        } 
        // next
        if nums[av-1] == av || nums[av-1] == 0 {
            nums[av-1] += nums[i]
            nums[i] = 0
            continue
        } 
        
        // linked list - chain of update
        cur := i
        next := abs(nums[i])-1
    
        for abs(nums[next]) != next + 1 || nums[next] != 0{ // update target 
            
            b := nums[next] // buff the next value
            nums[next] = nums[cur] // update the next with cur
            
            nums[cur] = 0
            cur = abs(b)

            next = abs(nums[abs(cur)])-1
        }
    }

    for i := len(nums)-1; i>=0; i --{
        if nums[i] == 0 {
            return i + 1
        }
    }
    return -1
}

func abs (n int) int {
    if n < 0 {
        return n * -1
    }
    return n
}
