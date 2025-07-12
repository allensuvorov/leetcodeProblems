func partitionString(s string) []string {
    seen := make(map[string]bool)
    segment := ""
    segments := make([]string, 0, len(seen))
    for _, c := range s {
        segment += string(c)
        if !seen[segment] {
            seen[segment] = true
            segments = append(segments, segment)
            segment = ""
        }
    }
    return segments
}
