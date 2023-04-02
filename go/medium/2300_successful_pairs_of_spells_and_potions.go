package medium

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
	// sort the potions
	sort.Ints(potions)
	res := make([]int, len(spells))

	for i := range spells {
		// use binary search
		res[i] = len(potions) - sort.Search(
			len(potions),
			func(j int) bool {
				return potions[j]*spells[i] >= int(success)
			},
		)
	}
	return res
}
