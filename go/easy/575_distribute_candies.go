func distributeCandies(candyType []int) int {
    n := len(candyType)
    types := make(map[int]struct{}, n) // set
    for _, v := range candyType {
        types[v] = struct{}{}
        if len(types) == n/2 {
            return n/2
        }
    }
    return len(types)
}
