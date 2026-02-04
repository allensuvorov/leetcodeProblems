func finalPrices(prices []int) []int {
    res := make([]int, len(prices))
    st := []int{}
    for i, v := range prices {
        res[i] = v
        for len(st) > 0 && prices[st[len(st)-1]] >= v {
            res[st[len(st)-1]] -= v
            st = st[:len(st)-1]
        }
        st = append(st, i)
    }
    return res
}
