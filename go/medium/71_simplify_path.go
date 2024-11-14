func simplifyPath(path string) string {
    stack := make([]string, 0)
    for _, name := range strings.Split(path, "/") {
        switch name {
        case "", ".": // do nothing
        case "..": // pop
            if len(stack) > 0 {
                stack = stack[:len(stack)-1]
            }
        default: // push
            stack = append(stack, name)
        }
    }
    return "/" + strings.Join(stack, "/")
}
