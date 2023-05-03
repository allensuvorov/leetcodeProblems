func arithmeticTriplets(nums []int, diff int) int {
    m := map[int]bool{}
    c := 0
    for _, v := range nums {
        m[v] = true
        if m[v-diff] && m[v-2*diff]{
            c++    
        }
    }
    return c
}
