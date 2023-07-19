func removeElement(nums []int, val int) int {
    l := 0
    r := len(nums) - 1
    for l <= r {
        if nums[l] == val {
            for l <= r {
                if nums[r] != val {
                    nums[l] = nums[r]
                    r--
                    l++
                    break
                } else {
                    r--
                }
            }
        } else {
            l++
        }
    }
    return l
}
