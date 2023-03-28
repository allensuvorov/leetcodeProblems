func repeatedCharacter(s string) byte {
    binary := 0
    for i := range s {
        shiftedBit := 1 << (s[i]-'a')
        if binary & shiftedBit == shiftedBit {
            return s[i]
        }
        binary ^= shiftedBit
    }
    return 0
}
