func findMin(nums []int) int {
    left, right := 0, len(nums)-1 // classic first match

	for left < right {
		mid := left + (right-left)/2 // classic first match
		if nums[mid] > nums[right] { // pop all left side higher then right (some of rotated)
			left = mid + 1
		} else if nums[mid] < nums[right] { // step to the smallest
			right = mid
		} else { // if equal to right pop that on right 
			right--
		}
	}

	return nums[left]
}
