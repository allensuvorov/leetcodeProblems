func finalString(s string) string {
    ans := make([]byte, 0, len(s))
    for i := range s {
        if s[i] == 'i' {
            reverse(ans)
        } else {
            ans = append(ans, s[i])
        }
    }
    return string(ans)
}

func reverse(s []byte){
    for l, r := 0, len(s) - 1; l < r; l, r = l+1, r-1 {
        s[l], s[r] = s[r], s[l]
    }
}
