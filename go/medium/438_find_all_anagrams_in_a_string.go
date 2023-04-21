package medium

func findAnagrams(s string, p string) []int {
    pCount := [26]int{}
    sCount := [26]int{}
    res:= []int{}
    
    if len(p) > len(s) {
        return []int{}
    }
    for i := range p {
        pCount[p[i]-'a']++
        sCount[s[i]-'a']++
    }

    if sCount == pCount {
         res = []int{0}
    }
    
    for i := 1; i <= len(s)-len(p); i++ {
        sCount[s[i-1]-'a']--
        sCount[s[i-1+len(p)]-'a']++
        fmt.Println(sCount)
        if sCount == pCount {
            res = append(res, i)
        }
    }
    return res
}
