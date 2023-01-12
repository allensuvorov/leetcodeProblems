func reverseVowels(s string) string {
    str := strings.Split(s, "")
    vs := map[string]bool{"a":true,"e":true,"i":true,"o":true,"u":true,"A":true,"E":true,"I":true,"O":true,"U":true}
    r := len(str)-1
    for l := 0; l < r; l++ {
        if _, ok := vs[str[l]]; ok {
            for ; r > l; r-- {
                if _, ok := vs[str[r]]; ok {
                    fmt.Println(str[l], l, str[r], r)
                    str[l], str[r] = str[r], str[l]
                    r--
                    break
                }
            }
        }
    }
    return strings.Join(str, "")
}
