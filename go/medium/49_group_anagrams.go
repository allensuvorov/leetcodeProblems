func groupAnagrams(strs []string) [][]string {
    groups := make(map[[26]int][]string)

    for _, str := range strs {
        counts := [26]int{}
        for _, v := range str {
            counts[v - 'a']++
        }
        groups[counts] = append(groups[counts], str)
    }
    res := make([][]string, 0, len(groups))
    for _, group := range groups {
        res = append(res, group)
    }
    return res
}
