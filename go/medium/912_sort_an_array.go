func sortArray(nums []int) []int {
    quickSort(nums, 0, len(nums) - 1)
    return nums
}

func quickSort(nums []int, l, r int) {
    if l < r {
        p := partition(nums, l, r)
        quickSort(nums, l, p)
        quickSort(nums, p + 1, r)
    }
}

func partition(nums []int, l, r int) int {
    pivot := nums[(l+r)/2]
    for l <= r {
        for nums[l] < pivot {
            l++
        }
        for nums[r] > pivot {
            r--
        }
        if l <= r {
            nums[l], nums[r] = nums[r], nums[l]
            l++
            r--
        }
    }
    return l - 1
}
