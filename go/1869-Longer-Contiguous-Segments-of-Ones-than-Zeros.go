func checkZeroOnes(s string) bool {
    var prev rune = -1
    cb := map[rune]int{}
    mx := map[rune]int{}
    for _, v := range s{
        cb[v]++
        if mx[v] < cb[v] {
            mx[v] = cb[v]
        }
        if v != prev {
            cb[prev] = 0
        }
        prev = v
    }
    return mx['1'] > mx['0']
}
