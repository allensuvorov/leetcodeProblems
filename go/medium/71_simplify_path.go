func simplifyPath(path string) string {
    stack := []string{}
    names := strings.Split(path, "/")
    
    for _, v := range names {
        switch v {
        case "", ".": // skip
        case "..": // pop
            if len(stack) > 0 {
                stack = stack[:len(stack)-1]
            }
        default:  // append
            stack = append(stack, v)
        }
    }

    return "/" + strings.Join(stack, "/")
}
