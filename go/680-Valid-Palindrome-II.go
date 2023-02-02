func validPalindrome(s string) bool {
    j := len(s)-1
    i := 0
    
    if ok, i, j := checkNext(i,j,s); !ok { 
        if ok, _, _ = checkNext (i+1, j, s); ok {
            return true
        }
        if ok, _, _ = checkNext(i, j-1, s); ok {
            return true
        }
        return false        
    }
    return true
}

func checkNext(i, j int, s string) (bool, int, int) {
    for i < j {
        fmt.Println(string(s[i]), string(s[j]))
        if s[i] != s[j] {
            return false, i, j
        }
        i++
        j--
    }
    return true, 0, 0
}
