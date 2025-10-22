func reverseVowels(s string) string {
    l, r := 0, len(s) - 1
    data := []byte(s)
    for l < r {
        if !isVowel(data[l]) {
            l++
            continue
        }
        if !isVowel(data[r]) {
            r--
            continue
        }

        data[l], data[r] = data[r], data[l]
        l++
        r--
    }
    return string(data)
}

func isVowel(c byte) bool {
    vowels := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
    for _, v := range vowels {
        if c == v {
            return true
        }
    }
    return false
}

// using slices.Contains
// - swap, with two pointers
// - search using a loop
// - check for a vowel
func reverseVowels(s string) string {
    vowels := []byte("aeiouAEIOU")
    temp := []byte(s)
    
    for l, r := 0, len(s) - 1; l < r; {
        if !slices.Contains(vowels, temp[l]) {
            l++
            continue
        }
        if !slices.Contains(vowels, temp[r]) {
            r-- 
            continue
        }
        
        temp[r], temp[l] = temp[l], temp[r]

        l++
        r--
    }
    return string(temp)
}
