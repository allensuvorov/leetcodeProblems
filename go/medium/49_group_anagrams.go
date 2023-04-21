package medium

func groupAnagrams(strs []string) [][]string {
    m := make(map[[26]int][]string)

    for i := range strs {
        cf := [26]int{} // char collection imprint
        for _, c := range strs[i] {
            cf[c-'a']++
        }
        m[cf] = append(m[cf], strs[i])
    }
    res := make([][]string, len(m))
    i := 0
    for k := range m {
        res[i] = m[k]
        i++
    }
    return res
}
