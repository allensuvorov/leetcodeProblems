func shortestCompletingWord(licensePlate string, words []string) string {
    lp := strings.ToLower(licensePlate)
    lpMap := map[rune]int{}
    for _, v := range lp {
        if v >= 97 && v <= 122 {
            lpMap[v]++
        }
    }
    res := ""
    min := 15
    for _, w := range words {
        wMap := map[rune]int{}
        for _, l := range w {
            wMap[l]++
        }

        includes := true
        for k, v := range lpMap {
            if wMap[k] < v {
                includes = false
                break 
            }
        }

        if includes {
            if len(w) < min {
                res = w
                min = len(w)
            }
        }
    }
    return res
}
