func checkIfPangram(sentence string) bool {
    cl := make([]int, 26)
    for i := range sentence {
        cl[sentence[i]-'a'] = 1
    }
    for _, v := range cl {
        if v == 0 {
            return false
        }
    }
    return true
}
