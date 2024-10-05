func trap(height []int) int {
    maxH := 0
    maxI := 0
    for i, v := range height {
        if maxH < v {
            maxH = v
            maxI = i
        }
    }

    water := 0 
    curMax := 0
    for i := 0; i < maxI; i++ {
        if curMax < height[i] {
            curMax = height[i]
        } else {
            water += curMax - height[i] 
        }
    }

    curMax = 0
    for i := len(height) - 1; i > maxI; i-- {
        if curMax < height[i] {
            curMax = height[i]
        } else {
            water += curMax - height[i] 
        }
    }

    return water
}
