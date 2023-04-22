package medium

func partitionLabels(s string) []int {
    res := []int{}
    size := 0
    end := 0
    lastIndexes := map[byte]int{}
    for i := range s {
        lastIndexes[s[i]]=i
    }
    for i := range s {
        size++
        if lastIndexes[s[i]] > end {
            end = lastIndexes[s[i]]
        }
        if i == end {
            res = append(res, size)
            size = 0
        }
    }
    return res   
}
