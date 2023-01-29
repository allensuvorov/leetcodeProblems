func numJewelsInStones(jewels string, stones string) int {
    jSet := map[rune]bool{}
    count := 0
    for _, v := range jewels {
        jSet[v] = true
    }
    for _, v := range stones {
        if jSet[v] {
            count++
        }
    }
    return count
}
