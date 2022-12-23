package main

func main() {
	nums1 := []int{1,2,2,1}
	nums2 := []int{2,2}
	intersect (nums1, nums2)
}

func intersect(nums1 []int, nums2 []int) []int {
	var num3 *[]int
	var newT int
	t, it, tHolder, target := getNewT(&nums1, &nums2)

	for {
		if foundPairAndRemoved (t, it, tHolder, target ) {
			*num3 = append(*num3, t)
			newT, it, tHolder, target = getNewT(&nums1, &nums2)
			if newT < t {

			}

		} else {
			(*tHolder)[it] = -1
			t, it, tHolder, target = getNewT(&nums1, &nums2)
		}

	}
	return *num3
}

func foundPairAndRemoved (t, it int, tHolder, target *[]int) bool {

	for i, v := range *target {
		if v < 0 {
			continue
		}
		if v == t {
			(*target)[i] = -1
			(*tHolder)[it] = -1
			return true
		}
	}
	return false
}

func getNewT(nums1, nums2 *[]int) (t, it int, tHolder, target *[]int) {
	min1, i1 := findMin(nums1)
	min2, i2 := findMin(nums2)
	if min1 > min2 {
		return min1, i1, nums1, nums2
	} else {
		return min2, i2, nums2, nums1
	}
}

func findMin(nums *[]int) (min, i int){
	min = 1000
	iMin := 0
	for i, v := range *nums {
		if v < 0 {
			continue
		}
		if v < min {
			min = v
			iMin = i
		}
	}
	return min, iMin
}
