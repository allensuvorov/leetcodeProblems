func convertToBase7(num int) string {
    ans := []byte{}
    signed := false
    if num < 0 {
        signed = true
        num = -num
    }
    if num == 0 {
        return "0"
    }
    for num != 0 {
        rem := num % 7
        byteRem := rem + '0'
        ans = append([]byte{byte(byteRem)}, ans...)
        num /= 7
    }
    if signed == true {
        return "-" + string(ans)
    }
    return string(ans)
}
