func searchInsert(nums []int, target int) int {

	l := len(nums) - 1 // last index
	ptr := l / 2

	for {
		if target == nums[ptr] {
			return ptr
		}
		if target > nums[ptr] && target < nums[ptr+1] {
			return ptr
		}

		if target < nums[ptr] && target > nums[ptr-1] {
			return ptr - 1
		}

		if target > nums[ptr] {
			ptr = (ptr + l) / 2 // mid to len
		} else {
			ptr = (ptr / 2) // 0 to mid
		}

		if ptr == l || ptr == 0 {
			if target > nums[ptr] {
				return ptr + 1
			}
			if target < nums[ptr] {
				return ptr - 1
			}
		}
		fmt.Println(ptr)

	}

	return len(nums)
}