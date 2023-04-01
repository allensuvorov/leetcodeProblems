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

// check this out, what a fine solution, high readability
func reverseVowels(s string) string {
	temp := []byte(s)
	i := 0
	j := len(s) - 1

	for i < j {
		if !isVowel(s[i]) {
			i++
			continue
		}

		if !isVowel(s[j]) {
			j--
			continue
		}

		temp[i], temp[j] = temp[j], temp[i]
		i++
		j--
	}

	return string(temp)
}

func isVowel(char byte) bool {
	return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' || char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'u' || char == 'U'
}
