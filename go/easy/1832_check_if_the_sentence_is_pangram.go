func checkIfPangram(sentence string) bool {
    cl := [26]int{}
    for _, v := range sentence {
        cl[v - 'a'] = 1
    }
    for _, v := range cl {
        if v == 0 {
            return false
        }
    }
    return true
}
