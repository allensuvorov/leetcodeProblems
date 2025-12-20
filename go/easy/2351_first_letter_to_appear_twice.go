// bitmask
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

// hash map
func repeatedCharacter(s string) byte {
    seen := [26]bool{}
    for i, v := range s {
        if seen[v - 'a'] {
            return s[i]
        }
        seen[v - 'a'] = true
    }
    return 0
}
