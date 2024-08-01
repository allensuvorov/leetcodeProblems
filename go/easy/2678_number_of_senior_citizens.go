func countSeniors(details []string) int {
    ans := 0
    for _, v := range details {
        age, err := strconv.Atoi(v[11:13])
        if err != nil {
            fmt.Println(err)
        }
        if age > 60 {
            ans++
        }
    }
    return ans
}
