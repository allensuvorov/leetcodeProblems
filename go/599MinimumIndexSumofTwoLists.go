func findRestaurant(list1 []string, list2 []string) []string {
	var m = make(map[string]int)
	var min_sum int = 2000
	var res []string

	for i := 0; i < len(list1); i++ {
		m[list1[i]] = i
	}

	for j := 0; j < len(list2) && j <= min_sum; j++ {
		if i, match := m[list2[j]]; match {
			if i+j < min_sum {
				res = nil
				res = append(res, list2[j])
				min_sum = i + j
			} else if i+j == min_sum {
				res = append(res, list2[j])
			}
		}
	}
	return res
}
