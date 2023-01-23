func findLHS(nums []int) int {

    hm := map[int]int{}
    for _, v := range nums {
        hm[v]++
    }
    max := 0
    for i := range hm {
        _, ok := hm[i+1]
        if ok && hm[i] + hm[i+1] > max {
            max = hm[i] + hm[i+1]
        }
    }
    return max
}
