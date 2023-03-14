func restoreString(s string, indices []int) string {
    bytesResult := make([]byte, len(s))
    for i, v := range indices {
        bytesResult[v] = s[i]
    }
    return string(bytesResult)
}
