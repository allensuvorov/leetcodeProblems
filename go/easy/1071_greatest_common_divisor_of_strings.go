func gcdOfStrings(str1 string, str2 string) string {
    res := ""
    for i := 0; i < len(str1) && i < len(str2); i++ {
        if str1[:i+1] == str2[:i+1] {
            pref := str1[:i+1]
            if len(str1) % len(pref) == 0 && len(str2) % len(pref) == 0 {
                isDevisor := true
                for i := 0; i <= len(str1) - len(pref); i += len(pref) {
                    if str1[i:i+len(pref)] != pref {
                        isDevisor = false
                        break
                    }
                }
                for i := 0; i <= len(str2) - len(pref); i += len(pref) {
                    if str2[i:i+len(pref)] != pref {
                        isDevisor = false
                        break
                    }
                }
                if isDevisor {
                    res = pref
                }
            }
        }
    }
    return res
}
