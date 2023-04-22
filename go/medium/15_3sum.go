package medium

func threeSum(nums []int) [][]int {    
    res := [][]int{}
    sort.Ints(nums)

    for i := 0; i < len(nums) && nums[i] <= 0; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        l := i + 1
        r := len(nums) - 1
        for l < r {
            sum := nums[i] + nums[l] + nums[r]
            switch {
            case sum < 0:
                l++
            case sum > 0:
                r--
            case sum == 0:
                res = append(res, []int{nums[i], nums[l], nums[r]})
                l++
                for nums[l] == nums[l-1] && l < r {
                    l++
                }
            }
        }
    }
    return res
}
