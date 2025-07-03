func kthCharacter(k int) byte {
    // simulation
    res := []byte{'a'}
    for len(res) < k {
        size := len(res)
        for i := range size {
            res = append(res, res[i]+1)
        }
    }
    return res[k-1]
}
