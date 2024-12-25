func compress(chars []byte) int {
    if len(chars) == 1 {
        return 1
    }
    lenth := 0
    count := 1
    for i := 0; i < len(chars); i++ {
        if i == len(chars)-1 || chars[i] != chars[i+1] {
            chars[lenth] = chars[i]
            lenth++
            if count > 1 {
                digits := []byte(strconv.Itoa(count))
                for i := range digits {
                    chars[i + lenth] = digits[i]
                }
                lenth += len(digits)
            }
            count = 1
        } else {
            count++
        }
    }
    return lenth
}
