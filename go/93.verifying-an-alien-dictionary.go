func isAlienSorted(words []string, order string) bool {
    if len(words) == 1 {
        return true
    }
    orderMap := map[byte]int{}
    for i := range order {
        orderMap[order[i]] = i
    }

    wordOneIsSmaller := func(w1, w2 string) bool {
        short := len(w1)
        if len(w1)>len(w2){
            short = len(w2)
        }
        for i := 0; i < short; i++ {
            if orderMap[w1[i]] < orderMap[w2[i]] {
                return true
            }
            if orderMap[w1[i]] > orderMap[w2[i]] {
                return false
            }
        }
        if len(w1)>len(w2){
            return false
        }
        return true
    } 

    for i := 0; i < len(words)-1; i++ {
        if !wordOneIsSmaller(words[i], words[i+1]) {
            return false
        }
    }
    
    return true
}
