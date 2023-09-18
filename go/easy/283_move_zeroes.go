func moveZeroes(nums []int)  {
    r := 0
    for l, v := range nums {
        if v == 0 {
            if r <= l {
                r = l + 1
            }
            for r < len(nums) && nums[r] == 0 {
                r++
            }
            if r == len(nums) {
                return
            } else {
                nums[l], nums[r] = nums[r], 0
            }
        }
    }
}
