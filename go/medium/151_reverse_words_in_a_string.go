// built-in methods solution
func reverseWords(s string) string {
    words := strings.Fields(s)
    slices.Reverse(words)
    return strings.Join(words, " ")
}

// no built-in methods solution - slice
func reverseWords(s string) string {
    words := []byte{}
    word := []byte{}
    for i := range s {
        if s[i] != ' ' {
            word = append(word, s[i])
        }
        if len(word) > 0 && (s[i] == ' ' || i == len(s)-1) {
            if len(words) > 0 {
                word = append(word, ' ')
            }
            words = append(word, words...)
            word = []byte{}
        }
    }
    return string(words)
}


// no built-in methods solution - string
func reverseWords(s string) string {
    var res string
    var word []rune
    for i, v := range s {
        if v != ' ' {
            word = append(word, v)
        } 
        if len(word) > 0 && (v == ' ' || i == len(s) - 1){    
            if len(res) > 0 {
                res = string(word) + " " + res
            } else {
                res = string(word)
            }
            word = []rune{}
        }
    }
    return res
}

// recurcive solution
func reverseWords(s string) string {
    if len(s) == 0 {
        return ""
    }

    // get first word
    var word string
    i := 0
    for i < len(s) {
        if s[i] != ' ' {
            word += string(s[i])
        } else if len(word) > 0 {
            break
        }
        i++
    }
    
    if len(word) == 0 {
        return ""
    }

    next := reverseWords(s[i:])
    
    if len(next) == 0 {
        return word
    }

    return next + " " + word
}
