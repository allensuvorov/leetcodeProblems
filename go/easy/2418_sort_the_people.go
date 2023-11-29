func sortPeople(names []string, heights []int) []string {
    nameHeight := make(map[int]string)
    for i := range names {
        nameHeight[heights[i]] = names[i] 
    }
    sort.Ints(heights)
    res := make([]string, 0, len(names))
    for i := len(heights)-1; i >= 0; i-- {
        res = append(res, nameHeight[heights[i]])
    } 
    return res
}
