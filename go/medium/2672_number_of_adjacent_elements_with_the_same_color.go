func colorTheArray(n int, queries [][]int) []int {
    colors := make([]int, n)
    count := 0 
    res := make([]int, len(queries))

    for i, v := range queries {
        index := v[0]
        color := v[1]
        // deduct
        count -=countPairs(colors, index)
        colors[index] = color
        // add
        count +=countPairs(colors, index)
        res[i] = count
    }
    return res
}

func countPairs(colors []int, index int) int {
    count := 0
    if colors[index] == 0 {
        return count
    }
    if index > 0 && colors[index] == colors[index - 1] {
        count++
    }
    if index < len(colors) - 1 && colors[index] == colors[index + 1] {
        count++
    }
    return count
}
