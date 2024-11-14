func simplifyPath(path string) string {
    path += "/"
    stack := make([]string, 0)
    name := make([]byte, 0)
    for i := range path {
        if path[i] == '/' {
            switch string(name) {
            case "", ".": // do nothing
            case "..": // pop
                if len(stack) > 0 {
                    stack = stack[:len(stack)-1]
                }
            default: // push
                stack = append(stack, string(name))
            }
            name = name[:0]
        } else {
            name = append(name, path[i])
        }
    }
    return "/" + strings.Join(stack, "/")
}
