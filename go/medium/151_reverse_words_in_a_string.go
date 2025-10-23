////////////////////////////////////////////////////////////////
// built-in methods solution
func reverseWords(s string) string {
    words := strings.Fields(s)
    slices.Reverse(words)
    return strings.Join(words, " ")
}
////////////////////////////////////////////////////////////////
// no built-in methods solution - linear run time
func reverseWords(s string) string {
    res := []byte{}
    n := len(s)
    // read s from end, parse each work and append to res
    l, r := n - 1, n - 1;
    for l >= 0 {
        if s[l] == ' ' {
            if l < r { // there is a word
                res = append(res, []byte(s[l+1:r+1])...)
                res = append(res, ' ')
            }
            r = l - 1 // r catches up with l, and jumps over
        }
        l--
    }

    if l < r {
        res = append(res, []byte(s[l+1:r+1])...)
        res = append(res, ' ')
    }

    return string(res)[:len(res)-1] // cut extra space at the end
}

////////////////////////////////////////////////////////////////
// no built-in methods solution - slice - O(n^2)
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
            words = append(word, words...) // prepending makes it O(n^2)
            word = []byte{}
        }
    }
    return string(words)
}

////////////////////////////////////////////////////////////////
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
////////////////////////////////////////////////////////////////
// recursive solution
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

////////////////////////////////////////
// in-place solution for a byte slice
func reverseWords(s string) string {
	data := []byte(s)

	// move extra spaces
	data = trimSpace(data)

	// reverse the entire string
	reverse(data)

	// reverse each word
    l, r := 0, 0
    for r < len(data) {
        if data[r] == ' ' {
            reverse(data[l:r])
            l = r + 1
        }
        r++
    }

    // last word
    reverse(data[l:r])

	return string(data)
}

func trimSpace(data []byte) []byte{
	// move spaces
	l, r := 0, 0
	firstSpace := true
	inWords := false
	for r < len(data) {
		if data[r] != ' ' {
			data[l], data[r] = data[r], data[l]
			l++
			firstSpace = true
			inWords = true
		} else if inWords && firstSpace {
			l++
			firstSpace = false
		}
		r++
	}
    
    data = data[:l]
    if data[l-1] == ' ' {
        data = data[:l-1]
    }
    return data
}

func reverse(data []byte) {
	l, r := 0, len(data)-1
	for l < r {
		data[l], data[r] = data[r], data[l]
        l++
        r--
	}
}

////////////////////////////////////////////////////////////////
