func findWinners(matches [][]int) [][]int {
    records := make([]int, 100001)

    for _, m := range matches {
        if records[m[0]] == 0 {
            records[m[0]] = 1
        }
        if records[m[1]] == 0 {
            records[m[1]] = 1
        }
        records[m[1]]++
    }

    answer := [][]int{{},{}}
    for i, rec := range records {
        if rec == 1 {
            answer[0] = append(answer[0], i)
        }
        if rec == 2 {
            answer[1] = append(answer[1], i)
        }
    }
    return answer
}
