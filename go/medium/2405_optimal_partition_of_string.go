func partitionString(s string) int {
    count := 0 
    bytes := make(map[byte]bool)
    
    for i := range s {
        if bytes[s[i]] {
            count ++
            bytes = map[byte]bool{}
        }
        bytes[s[i]] = true
    }
    
    return count+1
}
