func successfulPairs(spells []int, potions []int, success int64) []int {
    slices.Sort(potions)
    res := make([]int, len(spells))
    for i, spell := range spells {
        l, r := 0, len(potions) - 1
        for l <= r {
            m := l + (r-l)/2
            guess := int64(spell * potions[m])
            if guess >= success {
                r = m - 1
            } else {
                l = m + 1
            }
        }
        res[i] = len(potions) - l
    }
    return res
}
