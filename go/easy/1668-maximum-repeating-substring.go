func maxRepeating(sequence string, word string) int {
    max := 0
    l := len(word)
    for i := 0; i <= len(sequence)-l; i++ {
        count := 0
        for i <= len(sequence)-l && word == sequence[i:i+l] {
            count ++
            i += l
            fmt.Println(count, i, len(sequence))             
        }
        if count != 0 {
            i -= l-1
        }
        if max < count {
            max = count
        }
    }
    return max
}
