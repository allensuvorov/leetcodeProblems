package medium
// set implemented with array
func partitionString(s string) int {
    set := [26]bool{}
    count := 0
    for _, r := range s{
        if set[r-'a'] {
            count++
            set = [26]bool{}
        }
        set[r-'a'] = true
    }
    return count+1
}
// set implemented with hash map
func partitionString(s string) int {
    count := 1
    seen := make(map[byte]bool)
    for i := range s {
        if seen[s[i]] {
            count++
            seen = make(map[byte]bool)
        }
        seen[s[i]] = true    
    }
    return count
}
