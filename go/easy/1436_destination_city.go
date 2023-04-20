package easy

func destCity(paths [][]string) string {
    from := make(map[string]struct{})

    for p := range paths {
        from[paths[p][0]] = struct{}{}
    }

    for p := range paths {
        if _, ok := from[paths[p][1]]; !ok {
            return paths[p][1]
        }
    }
    return ""
}
