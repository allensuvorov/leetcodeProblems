package medium

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
	// sort the potions
	sort.Ints(potions)
	res := make([]int, len(spells))

	for i := range spells {
		// use binary search from standard library
		res[i] = len(potions) - sort.Search(
			len(potions),
			func(j int) bool {
				return potions[j]*spells[i] >= int(success)
			},
		)
	}
	return res
}

// implement binary search
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
    for l < r {
        m := l + (r-l)/2
        guess := int64(spell * potions[m])
        if guess >= success {
            r = m
        } else {
            l = m + 1
        }
    }
    if int64(spell * potions[l]) < success {
        return len(potions)
    }
    return l
}
