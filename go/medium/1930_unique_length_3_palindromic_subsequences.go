func countPalindromicSubsequence(s string) int {
    charFirstLastPos := map[byte][]int{}

    for i := range s {
        if pos, exists := charFirstLastPos[s[i]]; !exists {
            charFirstLastPos[s[i]] = []int{i,0}
        } else {
            pos[1] = i
        }
    }

    res := 0
    for _, pos := range charFirstLastPos {
        if pos[1] != 0 {
            charSet := map[byte]bool{} // count unique chars between positions
            for i := pos[0]+1; i < pos[1]; i++ {
                charSet[s[i]] = true
            }
            res += len(charSet)
        }
    }

    return res
}
    // one run - for each repeating char get positions: first and last
    // fill map of positions [first, last] (max len = 26)
    // go over the map between first and last and count unique chars 
