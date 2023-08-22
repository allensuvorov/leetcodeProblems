func addBinary(a string, b string) string {
    maxLen := len(a)
    if len(b) > len(a) {
        maxLen = len(b)
    }
    c := make([]byte, maxLen)
    var carry byte
    for i := 0; i < len(a) || i < len(b); i++ {
        var aBit, bBit, cBit byte
        if i < len(a) {
            aBit = a[len(a) - 1 - i] - '0'
        }
        if i < len(b) {
            bBit = b[len(b) - 1 - i] - '0'
        }
        cBit = carry ^ aBit ^ bBit // 0s and 1+1 parish, last 1 stays
        if carry + aBit + bBit > 1 {
            carry = 1
        } else {
            carry = 0
        }
        c[len(c) - 1 - i] = cBit + '0'
    }
    if carry == 1 {
        c = append([]byte{carry + '0'}, c...)
    }
    return string(c)
}
