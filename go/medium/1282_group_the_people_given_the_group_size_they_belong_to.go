func groupThePeople(groupSizes []int) [][]int {
    sizes := make(map[int][]int)
    groups := make([][]int, 0, len(sizes))

    for i, v := range groupSizes {
        sizes[v] = append(sizes[v], i)

        if len(sizes[v]) == v { // flush
            groups = append(groups, sizes[v])
            sizes[v] = []int{}
        }
    }
    return groups
}
