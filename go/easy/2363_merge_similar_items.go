func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
    weights := make([]int, 1001)
    for _, v := range items1 {
        weights[v[0]] += v[1]
    }
    for _, v := range items2 {
        weights[v[0]] += v[1]
    }

    res := [][]int{} 

    for i, v := range weights {
        if v > 0 {
            res = append(res, []int{i, v})
        }
    }
    return res
}
