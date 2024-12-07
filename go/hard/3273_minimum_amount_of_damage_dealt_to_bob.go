func minDamage(power int, damage []int, health []int) int64 {
	priority := make([]int, len(damage))
	for i := range priority {
		priority[i] = i
	}

	slices.SortFunc(priority, func(a, b int) int {

		timeToKillA := ceil(health[a], power)
		timeToKillB := ceil(health[b], power)

		if float64(damage[a])/float64(timeToKillA) < float64(damage[b])/float64(timeToKillB) {
			return -1
		}
		return 1
	})

	totalDPS := 0
	for _, v := range damage {
		totalDPS += v
	}

	minDamage := 0

	for i := len(priority) - 1; i >= 0; i-- {

		timeToKill := ceil(health[priority[i]], power)
		minDamage += timeToKill * totalDPS // take damage
		totalDPS -= damage[priority[i]] // kill
	}

	return int64(minDamage)
}

func ceil(a, b int) int {
	ceil := a / b

	if a%b > 0 {
		ceil++
	}
    return ceil
}
