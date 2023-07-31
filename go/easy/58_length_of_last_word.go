solution 1
func lengthOfLastWord(s string) int {
	var res, i int
	// for i = len(s)-1; i >= 0; i-- {
	//     if s[i] != ' ' {
	//         res++
	//     } else if res != 0 {
	//         break
	//     }
	// }
	return res
}

solution 2
func lengthOfLastWord(s string) int {
    res := 0
    i := len(s)-1
    for ; s[i] == ' ' && i > 0; i-- {
    }
    
    for ; s[i] != ' ' && i > 0; i-- {
        res++
    } 
    return res
}

solution 3
func lengthOfLastWord(s string) int {
    inWord := false
    ans := 1
    for i := len(s) - 1; i >=0; i-- {
        if !inWord {
            if s[i] != ' ' {
                inWord = true
            }
        } else {
            if s[i] != ' ' {
                ans++
            } else {
                return ans
            }
        }
    }
    return ans
}
