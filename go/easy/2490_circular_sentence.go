import "strings"
func isCircularSentence(sentence string) bool {
    words := strings.Fields(sentence)
    for i, word := range words {
        last := len(word)-1
        cur := word[last]
        var next byte
        if i == len(words) - 1 {
            next = words[0][0]
        } else {
            next = words[i + 1][0]
        }
        if cur != next {
            return false
        }
        
    }
    return true
}
