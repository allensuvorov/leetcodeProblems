func toGoatLatin(sentence string) string {
    words := strings.Split(sentence, " ")
    for i := range words{
        word := []byte(words[i])

        switch word[0] {
        case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
            word = append(word, 'm', 'a')
        default:
            word = append(word[1:], word[0], 'm', 'a')
        }

        word = append(word, bytes.Repeat([]byte{'a'}, i+1)...)
        words[i] = string(word)
    }
    return strings.Join(words, " ")
}
