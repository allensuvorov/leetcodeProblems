func nextPermutation(nums []int)  {
    n := len(nums)
    if n == 1 {
        return
    }

    l, r := n - 2, n - 1
    for l >= 0 && nums[l] >= nums[r] {
        l--
        r--
    }

    
    for i := r; l >= 0 && i < n; i++ {
        if nums[i] > nums[l] {
            r = i
        }
    }

    if r > 0 {
        nums[l], nums[r] = nums[r], nums[l]
    }

    slices.Sort(nums[l+1:]) // O(n log n)
}

    // find most right inc - swap with next val (smalles of the bigger ones)
    // , and inc sort from r val
    // if no inc - inc sort

