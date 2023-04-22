func maxArea(height []int) int {
    max := 0
    l, r := 0, len(height)-1
    for l < r {
        space := min(height[l], height[r]) * (r - l)
        if max < space {
            max = space
        }
        switch {
        case height[l] <= height[r]: l++
        case height[l] > height[r]: r--
        } 
    }
    return max
}

func min (a,b int) int{
    if a > b {
        return b
    } else {
        return a
    }
}
