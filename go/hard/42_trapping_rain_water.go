func trap(height []int) int {
    // find max
    max := 0
    maxI := 0
    for i, v := range height{
        if v > max {
            max = v
            maxI = i
        }
    }

    // count water
    res := 0

    max = 0
    for l := 0; l < maxI; l ++ {
        if max < height[l] {
            max = height[l]
        } else {
            res +=(max - height[l])
        }

    }
    
    max = 0
    for r := len(height)-1; r > maxI; r -- {
        if max < height[r] {
            max = height[r]
        } else {
            res +=(max - height[r])
        }
    }
    return res
}
