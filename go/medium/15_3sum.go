func threeSum(nums []int) [][]int {
	n := len(nums)
	slices.Sort(nums)
	tripletSet := map[[3]int]bool{}
    triplets := [][]int{}

	for i := range nums {
		j := i + 1
		k := len(nums) - 1
		for j < k && j < n-1 && k > i {
			sum := nums[i] + nums[j] + nums[k]
			if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			} else {
				tripletSet[[3]int{nums[i], nums[j], nums[k]}] = true
				j++
                k--
			}
		}
	}
    for triplet := range tripletSet{
        triplets = append(triplets, triplet[:])
    }
	return triplets
}
