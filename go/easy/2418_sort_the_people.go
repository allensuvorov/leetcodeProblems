func sortPeople(names []string, heights []int) []string {
    m := make(map[int]string, len(names))
    for i, v := range names {
        m[heights[i]] = v
    }
    sort.SliceStable(heights, func(i, j int) bool { return heights[i] > heights[j]} )
    
    for i, v := range heights {
        names[i] = m[v]
    }
    return names
}
