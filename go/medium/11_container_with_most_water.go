func maxArea(height []int) int {
    cap := 0
    for l, r := 0, len(height) - 1; l < r; {
        cap = max(cap, (r - l) * min(height[r], height[l]))
        if height[l] < height[r] {
            l++
        } else {
            r--
        }
    }
    return cap
}
