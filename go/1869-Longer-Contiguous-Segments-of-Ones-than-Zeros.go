func checkZeroOnes(s string) bool {
    var prev rune = -1
    cb := map[rune]int{}
    mx := map[rune]int{}
    for _, v := range s{
        cb[v]++
        if v != prev{
            if mx[prev] < cb[prev] {
                mx[prev] = cb[prev]
            }
            cb[prev] = 0
        }
        prev = v
    }
    lr := rune(s[len(s)-1])
        if mx[lr] < cb[lr] {
            mx[lr] = cb[lr]
        }
    return mx['1'] > mx['0']
}
