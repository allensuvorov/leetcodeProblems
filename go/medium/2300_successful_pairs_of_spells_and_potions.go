func successfulPairs(spells []int, potions []int, success int64) []int {
    slices.Sort(potions)
    res := make([]int, len(spells))
    for i, spell := range spells {
        res[i] = len(potions) - binarySearch(potions, spell, success)
    }
    return res
}

func binarySearch(potions []int, spell int, success int64) int {
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
    return l
}
