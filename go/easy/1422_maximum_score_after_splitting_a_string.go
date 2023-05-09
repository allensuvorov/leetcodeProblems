func maxScore(s string) int {
    zeros, ones := 0, 0
    for i := range s {
        if s[i] == '1' {
            ones++
        }
    }
    max := 0
    for i := 0; i < len(s) - 1; i ++ {
        if s[i] == '1' {
            ones--
        } else {
            zeros++
        }

        if max < ones+zeros{
            max = ones+zeros
        }
    }
    return max
} 
