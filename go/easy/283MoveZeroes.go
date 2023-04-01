package leetcode

func moveZeroes(nums []int) {

	// go over nums
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			// go from after zero
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}
	}
}
