func findRepeatedDnaSequences(s string) []string {
    status := make(map[string]int) // 0 - not seen, 1 - seen, 2 - add to result
    result := []string{}
    for i := 0; i <= (len(s) - 10); i++ {
        substring := s[i:i+10]
        switch status[substring]{
        case 0:
            status[substring] = 1
        case 1:
            result = append(result, substring)
            status[substring] = 2
        }
    }
    return result
}
