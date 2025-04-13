// time Big O(30 * n)
func maxDistance(position []int, m int) int {
    slices.Sort(position)
    l, r := 1, position[len(position) - 1] - position[0]
    ans := 0
    for l <= r {
        guess := l + (r - l)/2
        if canPlace(position, m, guess) {
            l = guess + 1
            ans = guess
        } else {
            r = guess - 1
        }
    }
    return ans
}

// time Big O(n)
func canPlace(position []int, m int, force int) bool {
    prev := position[0]
    m--
    for _, v := range position {
        if v - prev >= force {
            m--
            prev = v
        } 
    }
    return m <= 0
}
