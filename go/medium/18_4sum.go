func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    res := [][]int{}
    for a := range nums{
        if a != 0 && nums[a] == nums[a-1] {
            continue
        }
        for b := a + 1; b < len(nums)-2; b++ {
            if b != a + 1 && nums[b] == nums[b-1] {
                continue
            }
            l := b + 1
            r := len(nums)-1
            for l < r {
                sum := nums[a] + nums[b] + nums[l] + nums[r]
                switch {
                case sum > target:
                    r--
                case sum < target:
                    l++
                case sum == target:
                    res = append(res, []int{nums[a], nums[b], nums[l], nums[r]})
                    l++
                    for nums[l] == nums[l-1] && l < r{
                        l++
                    }
                }
            }
        }
    }
    return res
}
