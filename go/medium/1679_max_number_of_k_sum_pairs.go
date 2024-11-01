func maxOperations(nums []int, k int) int {
    fm := map[int]int{}
    for _, v := range nums {
        fm[v]++
    }
    ans := 0
    for num, v := range fm {   
        if num * 2 == k {
            ans += v / 2 
        } else {
            temp := min(v, fm[k - num])
            ans += temp
            fm[num] -= temp
            fm[k - num] -= temp
        }
    }
    return ans
}
