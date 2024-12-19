// built-in methods solution
func reverseWords(s string) string {
    words := strings.Fields(s)
    slices.Reverse(words)
    return strings.Join(words, " ")
}

// no built-in methods solution
func reverseWords(s string) string {
    var res, word string
    for i, v := range s {
        if v != ' ' {
            word += string(v)
        } 
        if (v == ' ' || i == len(s)-1) && len(word) > 0 {
            if len(res) == 0 {
                res = word
            } else {
                res = word + " " + res
            }
            word = ""

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
