func separateDigits(nums []int) []int {
    stack := make([]int, 0, len(nums))
    for i := len(nums)-1; i >= 0; i-- {
        for tmp := nums[i]; tmp > 0; tmp /= 10 {
            stack = append(stack, tmp % 10)
        }
    }
    ans := make([]int, len(stack))
    for i := range ans {
        top := len(stack) - 1
        ans[i] = stack[top]
        stack = stack[:top]
    }
    return ans
}
