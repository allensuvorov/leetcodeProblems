func hasGroupsSizeX(deck []int) bool {
    if len(deck) == 1{
        return false
    }
    hm := make(map[int]int)
    for i := range deck{
        hm[deck[i]]++
    }
    
    min := 10000
    for i := range hm{
        if hm[i] < min{
            min = hm[i]
        }
    }

    if min == 1 {
        return false
    }

    for x := 2; x <= min; x ++ {
        ok := true
        for i := range hm{
            if hm[i] % x != 0 {
                ok = false
                break
            }
        }
        if ok {
            return true
        }
    }
    return false
}
