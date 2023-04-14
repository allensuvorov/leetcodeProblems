package easy

func heightChecker(heights []int) int {
    heightCount := [101]int{}
    
    for _, height := range heights {
        heightCount[height]++
    }

    p := 0
    res := 0

    for h, c := range heightCount {
        for c > 0 {
            if heights[p] != h {
                res++
            }
            p++
            c--
        }
    }

    return res
}
