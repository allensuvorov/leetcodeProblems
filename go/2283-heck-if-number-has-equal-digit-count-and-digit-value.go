func digitCount(num string) bool {
    hm := map[int]int{}
    for i := range num {
        v := int(num[i] - '0')
        hm[i] += v
        hm[v] --
        if hm[v] == 0 {
            delete(hm, v)
        }
        if hm[i] == 0 {
            delete(hm, i)
        }
    }
    return len(hm) == 0
}
