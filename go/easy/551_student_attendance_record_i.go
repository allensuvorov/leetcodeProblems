func checkRecord(s string) bool {
    a := 0 // absent < 2 total 
    l := 0 // late < 3 in a row

    for i := range s {
        if s[i] =='L' {
            l++
            if l > 2 {
                return false
            }
        } else {
            l = 0
            if s[i] == 'A' {
                if a > 0 {
                    return false
                } else {
                    a++
                }
            } 
        }
    }
    return true
}
