func finalString(s string) string {
    ans := make([]byte, 0, len(s))
    needToReverse := -1
    for i := range s {
        if s[i] == 'i' {
            needToReverse *= -1
        } else {
            if needToReverse == 1 {
                reverse(ans)
                needToReverse = -1    
            }
            ans = append(ans, s[i])
        }
    }
    if needToReverse == 1 {
        reverse(ans)
    }
    return string(ans)
}

func reverse(s []byte){
    for l, r := 0, len(s) - 1; l < r; l, r = l+1, r-1 {
        s[l], s[r] = s[r], s[l]
    }
}
