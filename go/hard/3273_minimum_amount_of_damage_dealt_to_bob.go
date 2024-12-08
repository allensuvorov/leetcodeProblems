func minDamage(power int, damage []int, health []int) int64 {
	priority := make([]int, len(damage))
	for i := range priority {
		priority[i] = i
	}

	slices.SortFunc(priority, func(a, b int) int {
		timeA := ceil(health[a], power)
		timeB := ceil(health[b], power)
		return compare(damage[a], timeA, damage[b], timeB)
	})

	totalDPS := 0
	for _, v := range damage {
		totalDPS += v
	}

	var minDamage int64 = 0
	for i := len(priority) - 1; i >= 0; i-- {
		time := ceil(health[priority[i]], power)
		minDamage += int64(time * totalDPS) // take damage
		totalDPS -= damage[priority[i]] // kill
	}

	return int64(minDamage)
}

func compare(damageA, timeA, damageB, timeB int) int {
    if damageA * timeB < damageB * timeA {
        return - 1
    }
    return 1
}

func ceil(a, b int) int {
    return (a + b - 1)/b 
}
