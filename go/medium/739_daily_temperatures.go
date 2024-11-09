func dailyTemperatures(temperatures []int) []int {
    ans := make([]int, len(temperatures))
    stack := []int{}

    for i := range temperatures {
        for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack) - 1]]{
            top := len(stack) - 1
            ans[stack[top]] = i - stack[top]
            stack = stack[:top]
        }
        stack = append(stack, i)
    }
    return ans
}
