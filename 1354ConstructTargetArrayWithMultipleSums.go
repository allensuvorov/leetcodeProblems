// first idea - start from the end result (reverse engineering)
func isPossible(target []int) bool {
	max := 0
	sum := 0
	maxI := 0
	rest := 0 // rest of array, excluding max

	for {
		sum = 0
		max = 0
		for i, num := range target {
			sum += num
			if num > max {
				max = num
				maxI = i
			}
		}
		//fmt.Println(target, sum, max, maxI)
		if max == 1 {
			return true
		}
		rest = sum - max
		if max < rest+1 {
			return false
		} else {
			target[maxI] = max % rest // replace max with remainder
		}
	}
}
