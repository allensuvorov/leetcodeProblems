func reversePrefix(word string, ch byte) string {
    var l int
    var chExists bool
    for l = range word {
        if word[l] == ch {
            chExists = true
            break
        }
    }
    if !chExists {
        return word
    }
    stack := []byte{}
    for r := l; r >=0; r-- {
        stack = append(stack, word[r])
    }
    return string(stack) + word[l+1:]
}
