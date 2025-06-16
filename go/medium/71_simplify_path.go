func simplifyPath(path string) string {
    stack := []string{}
    for _, v := range strings.Split(path, "/") {
        switch {
        case v == "..": 
            // pop
            if len(stack) > 0 {
                stack = stack[:len(stack) - 1]
            }
        case len(v) > 0 && v != ".":
            // append
            stack = append(stack, v)
        }
    }
    return "/" + strings.Join(stack, "/")
}
