func removeSubfolders(folder []string) []string {
    // hash map -> Big O(l * l * n)
    prefix := make(map[string]bool)

    // soft by path len in inc order, to start from shorter ones
    slices.SortFunc(folder, func(a, b string) int {
        return len(a) - len(b)
    })

    for _, path := range folder {
        // if prefix is in the map, skip, else - add to the map
        hasPrefix := false
        for i := range path {
            if path[i] == '/' && prefix[path[:i]] {
                hasPrefix = true
                break
            }
        }
        if !hasPrefix {
            prefix[path] = true
        }
    }
    res := make([]string, 0, len(prefix))
    for path := range prefix {
        res = append(res, path)
    }
    return res
}
