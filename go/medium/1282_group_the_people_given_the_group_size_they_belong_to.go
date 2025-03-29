func groupThePeople(groupSizes []int) [][]int {
    sizes := make(map[int][]int)
    for i, v := range groupSizes {
        sizes[v] = append(sizes[v], i)
    }

    groups := make([][]int, 0, len(sizes))

    for size, items := range sizes {
        group := make([]int, 0, size)

        for i := range items {
            group = append(group, items[i])
            if len(group) == size { // flush
                groups = append(groups, group)
                group = make([]int, 0, size)
            }
        }
    }
    return groups
}
