func digitSum(s string, k int) string {
    for k < len(s){
        s = round(s, k)
    }
    return s
}

func round(s string, k int) string {
    count := 0
    sum := 0
    var sb strings.Builder
    for i := range s {
        count ++
        sum += int(s[i]-'0')
        if count == k || i == len(s)-1 {
            sb.WriteString(fmt.Sprint(sum))
            count = 0
            sum = 0
        }
    }
    return sb.String()
}
