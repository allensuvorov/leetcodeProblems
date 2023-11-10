func arrayRankTransform(arr []int) []int {
    // sort, then rank map: O (n log n)
    sorted := make([]int, len(arr))
    copy(sorted, arr)
    sort.Ints(sorted)
    rm := map[int]int{}
    rank := 1
    for _, v := range sorted {
        if _, ok := rm[v]; !ok {
            rm[v] = rank
            rank++
        }
    }
    for i := range arr {
        arr[i] = rm[arr[i]]
    }
    return arr
}
