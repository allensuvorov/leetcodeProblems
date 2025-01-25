func lexicographicallySmallestArray(nums []int, limit int) []int {
    // sort, keeping the indices
    // locate cluster, and sort it indices set
    // put values back into those spots

    items := make([][]int, len(nums))
    for i, num := range nums {
        items[i] = []int{num, i}
    }

    slices.SortFunc(items, func(a, b []int) int { return a[0]-b[0] })

    
    indices := []int{items[0][1]}
    values := []int{items[0][0]}

    res := make([]int, len(nums))

    for i := 1; i < len(items); i++ {
        if items[i][0] - items[i-1][0] <= limit {
            values = append(values, items[i][0])
            indices = append(indices, items[i][1])
            continue
        }

        slices.Sort(indices) 
        for j, index := range indices {
            res[index] = values[j]
        }
        

        indices = []int{items[i][1]}
        values = []int{items[i][0]}
    }

    slices.Sort(indices) 
    for j, index := range indices {
        res[index] = values[j]
    }

    return res
}
