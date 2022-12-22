func reverseString(s []byte)  {
    j := len(s)-1
    for i := 0; j - i >= 1; i++ {
        s[i], s[j] = s[j], s[i]
        j = len(s) - 2 - i
    }
}
