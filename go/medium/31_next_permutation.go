func nextPermutation(nums []int)  {
    n := len(nums)
    if n == 1 {
        return
    }

    i := n - 2
    for i >= 0 && nums[i] >= nums[i+1] {
        i--
    }
    
    k := 0
    for j := i+1; i >= 0 && j < n; j++ {
        if nums[j] > nums[i] {
            k = j
        }
    }

    if k > 0 {
        nums[i], nums[k] = nums[k], nums[i]
    }

    l := i + 1
    r := n - 1
    for l < r {
        nums[l], nums[r] = nums[r], nums[l]
        l++
        r--
    }
}

    // find most right inc - swap with next val (smalles of the bigger ones)
    // , and inc sort from r val
    // if no inc - inc sort

