func maxDistance(arrays [][]int) int {
    maxDist := 0
    minSoFar, maxSoFar := arrays[0][0], arrays[0][len(arrays[0])-1]
    for i := 1; i < len(arrays); i++ {
        left, right := arrays[i][0], arrays[i][len(arrays[i])-1]
        maxDist = max(maxDist, max(right - minSoFar, maxSoFar - left))
        minSoFar = min(minSoFar, left)
        maxSoFar = max(maxSoFar, right)
    }
    return maxDist
}
