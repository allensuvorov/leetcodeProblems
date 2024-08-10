func validStrings(n int) []string {
    var res = []string{}
    var dfs func(now byte, path []byte)
    dfs = func(now byte, path []byte) {
        if len(path) == n {
            res = append(res, string(path))
            return
        }
        if now == '1' {
            dfs('0', append(path, '0'))
        }
        dfs('1', append(path, '1'))
    }
    dfs('0', []byte{'0'})
    dfs('1', []byte{'1'})
    return res
}


//         0           1
//          \        /   \
//           1      0     1
