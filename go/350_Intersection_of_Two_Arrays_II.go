package main

func main() {
	nums1 := []int{1,2,2,1}
	nums2 := []int{2,2}
	intersect (nums1, nums2)
}

func intersect(nums1 []int, nums2 []int) []int {
    m := make(map[int]int)
    for _, v := range nums1 {
        m[v]++
    }
    
    nums3 := []int{}
    for _, v := range nums2 {
        _, ok := m[v]
        if ok {
            if m[v] > 0 {
                nums3 = append (nums3, v)
                m[v]--
            }
            if m[v] == 0 {
                delete(m, v)
            }
        }
    }
    return nums3
}
