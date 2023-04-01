func rearrangeCharacters(s string, target string) int {
    tm := map[rune]int{}
    sm := map[rune]int{}

    for _, v := range target {
        tm[v]++
    }

    for _, v := range s {
        _, ok := tm[v]
        if ok {
            sm[v]++
        }
    }
    min := 100
    for k := range tm {
        times := sm[k]/tm[k]
        if min > times {
            min = times
        }
    }
    return min
}
