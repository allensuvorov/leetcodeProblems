func largestGoodInteger(num string) string {
    var res string
    var max byte
    for i := 0; i <= len(num) - 3; i ++ {
        if num[i] == num[i+1] && num[i] == num[i+2] {
            if max < num[i] {
                max = num[i]
                res = string([]byte{max, max, max})
            }
        }
    }
    return res
}
