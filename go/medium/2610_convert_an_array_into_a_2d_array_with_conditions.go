func findMatrix(nums []int) [][]int {
    count := make(map[int]int)
    maxCount := 0

    for _, v := range nums {
        count[v]++
        maxCount = max(maxCount, count[v])
    }

    res := make([][]int, maxCount)
    
    for num := range count {
        for i := range count[num] {
            res[i] = append(res[i], num)
        }
    }
    return res
}
