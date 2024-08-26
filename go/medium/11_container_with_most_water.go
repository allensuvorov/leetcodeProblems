func maxArea(height []int) int {
    ans := 0
    l, r := 0, len(height) - 1

    for l < r {
        area := min(height[l], height[r]) * (r - l)
        ans = max(ans, area)
        if height[l] <= height[r] {
            l++
        } else {
            r--
        }
    }

    return ans
}
