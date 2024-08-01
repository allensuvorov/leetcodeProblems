func countSeniors(details []string) int {
    ans := 0
    for _, v := range details {
        age, _ := strconv.Atoi(v[11:13])
        if age > 60 {
            ans++
        }
    }
    return ans
}
