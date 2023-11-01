func largestGoodInteger(num string) string {
    var res string
    for i := 0; i <= len(num) - 3; i ++ {
        if num[i] == num[i+1] && num[i] == num[i+2] && num[i:i+3] > res {
            res = num[i:i+3]
        }
    }
    return res
}
