func findLongestChain(pairs [][]int) int {
    var count = len(pairs)
    sort.Slice(pairs, func(i, j int) bool {
        return pairs[i][1] < pairs[j][1]
    })
    fmt.Println(pairs)

    maxEnd := pairs[0][1]
    pairs = pairs[1:]
    for _, pair := range pairs {
        if (maxEnd >= pair[0]) {
            count -= 1
        } else {
            maxEnd = pair[1]
        }
    }
    return count
}
