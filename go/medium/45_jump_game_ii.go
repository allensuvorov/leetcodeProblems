func jump(nums []int) int {
    res := 0 
    l, r := 0, 0
    for r < len(nums) - 1 {
        farthest := 0
        for i := l; i <= r; i++ {
            if farthest < i + nums[i] {
                farthest = i + nums[i]
            }
        }
        l = r + 1
        r = farthest
        res++
    }
    return res
}
