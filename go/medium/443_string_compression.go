func compress(chars []byte) int {
	n := len(chars)
    i := 0
    res := 0
	
    for i < n {
	    // read a group
        groupLength := 1
        for i + groupLength < n && chars[i + groupLength] == chars[i] {
            groupLength++
        }

        // write the group
        chars[res] = chars[i]
        res++
        if groupLength > 1 {
            strRepr := strconv.Itoa(groupLength)
            for i := range strRepr {
                chars[res] = strRepr[i]
                res++
            }
        }
        i += groupLength
    }   
    return res
}
