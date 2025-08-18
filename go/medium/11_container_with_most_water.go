func maxArea(height []int) int {
    result := 0
    l := 0
    r := len(height) - 1
    for l < r {
        area := min(height[l], height[r]) * (r - l)
        result = max(result, area)
        if height[l] < height[r] {
            l++
        } else {
            r--
        }
    }
    return result
}
