func longestCommonPrefix(strs []string) string {
    // iterate array
    // check 1st, 2nd ... letter in each string, is it the same?
        // if different - return result, 
        // else append letter to result
    //
    // capture letters one by one 

    var (
        res []byte
        b byte
        bi int // byte index
    )
    
    for {
        if bi > len(strs[0])-1 {
                return string (res)
            }
        b = strs[0][bi]
        
        for i:=1; i<len(strs); i++ {
            if bi > len(strs[i])-1 {
                return string (res)
            }
            if b != strs[i][bi]{
                return string (res)
            }
        }
        res = append(res, b)
        bi++
    }
}

