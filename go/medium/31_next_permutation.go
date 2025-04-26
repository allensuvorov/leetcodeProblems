func nextPermutation(nums []int)  {
    i := len(nums) - 2
    for i >= 0 && nums[i] >= nums[i+1] {
        i--
    }
    
    if i >= 0 {
        j := len(nums) - 1
        for nums[j] <= nums[i] {
            j--
        }
        nums[i], nums[j] = nums[j], nums[i]
    }

    l := i + 1
    r := len(nums) - 1
    for l < r {
        nums[l], nums[r] = nums[r], nums[l]
        l++
        r--
    }
}

    // find most right inc - swap with next val (smallest of the bigger ones)
    // and inc sort from r val
    // if no inc - inc sort via reversing

