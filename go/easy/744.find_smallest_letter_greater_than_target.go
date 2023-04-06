func nextGreatestLetter(letters []byte, target byte) byte {
    l := 0
    r := len(letters)-1
    m := -1
    for l < r {
        m = (l+r)/2
        guess := letters[m]
        fmt.Println(string(guess))
        if guess > target {
            r = m
        } else {
            l = m + 1
        }
    }
    if letters[l] <= target {
        return letters[0]
    }
    return letters[l]
}
