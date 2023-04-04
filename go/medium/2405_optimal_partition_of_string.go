package medium

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
