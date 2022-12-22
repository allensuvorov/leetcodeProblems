func reverseString(s []byte)  {
    if len(s) > 0 {
        for b, e := 0, len(s)-1; e > b; {
            s[b], s[e] = s[e], s[b]
            b++
            e--
        }
    }
}
