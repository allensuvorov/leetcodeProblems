type FindSumPairs struct {
    nums1 []int
    nums2 []int
    counts1 map[int]int
    counts2 map[int]int
}


func Constructor(nums1 []int, nums2 []int) FindSumPairs {
    counts1 := make(map[int]int)
    counts2 := make(map[int]int)
  
    for _, v := range nums1 {
        counts1[v]++
    }
    for _, v := range nums2 {
        counts2[v]++
    }

    return FindSumPairs{
        nums1: nums1,
        nums2: nums2,
        counts1: counts1,
        counts2: counts2,
    }
}


func (this *FindSumPairs) Add(index int, val int)  {
    // update counts2
    old := this.nums2[index]
    new := this.nums2[index] + val
    this.counts2[old]--
    this.counts2[new]++

    // update nums2
    this.nums2[index] += val
}


func (this *FindSumPairs) Count(tot int) int {
    totalCount := 0
    for term1, count1 := range this.counts1 {
        term2 := tot - term1
        count2 := this.counts2[term2]
         totalCount += count1 * count2
    }
    return totalCount
}


/**
 * Your FindSumPairs object will be instantiated and called as such:
 * obj := Constructor(nums1, nums2);
 * obj.Add(index,val);
 * param_2 := obj.Count(tot);
 */
