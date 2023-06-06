func nextGreaterElement(nums1 []int, nums2 []int) []int {
    ans := make([]int, len(nums1))
    m := make(map[int]int)

    for i := range nums2 {
        m[nums2[i]] = i
    }

    for i, v:= range nums1 {    
        ans[i] = findNext(nums2, m[v], v)
    }
    return ans
}

func findNext(arr []int, i, n int) int {
    for j := i; j < len(arr); j ++ {
        if arr[j] > n {
            return arr[j]
        }
    }
    return - 1
}
