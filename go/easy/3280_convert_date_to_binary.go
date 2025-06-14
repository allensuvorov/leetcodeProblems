func convertDateToBinary(date string) string {
    res := ""
    for _, v := range strings.Split(date, "-") {
        if len(res) > 0 {
            res += "-"
        }
        num, _ := strconv.Atoi(v)
        res += fmt.Sprintf("%b", num)
    }
    return res
}
