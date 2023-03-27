func hasGroupsSizeX(deck []int) bool {
    if len(deck) == 1{
        return false
    }
    hm := make(map[int]int)
    for i := range deck{
        hm[deck[i]]++
    }
    for x := 2; x < 10000/2; x ++ {
        sameDenom := true
        for i := range hm{
            if hm[i] % x != 0 {
                sameDenom = false
                break
            }
        }
        if sameDenom {
            return true
        }
    }
    return false
}
