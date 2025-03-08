func minimumRecolors(blocks string, k int) int {
    minRec := k
    whiteCount := 0
    for r := range blocks {
        if blocks[r] == 'W' {
            whiteCount++
        }
        if r - k >= 0  { // we have a tail, we need to cut
            if blocks[r - k] == 'W' {
                whiteCount--
            }
        }
        if r + 1 >= k { // we have a window
            minRec = min(minRec, whiteCount)
        }
    }
    return minRec
}
