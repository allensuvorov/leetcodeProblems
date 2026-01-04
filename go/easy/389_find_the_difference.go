func findTheDifference(s string, t string) byte {
    var sum byte
    for i := range s {
        sum += t[i]
        sum -= s[i]
    }
    sum += t[len(t)-1]
    return sum
}

// bitwise solution
func findTheDifference(s string, t string) byte {
    var addedLetter byte
    addedLetter ^= t[len(t) - 1]
    for i := range s {
        addedLetter^=s[i]
        addedLetter^=t[i]
    }
    return addedLetter
}

// concurrent solution with idiomatic signaling
func findTheDifference(s string, t string) byte {
    var bitMask1 byte
    c := make(chan int)
    go func() {
        for i := range s {
            bitMask1 ^=s[i]
        }
        c <- 1
    }()
    
    var bitMask2 byte
    for i := range t {
        bitMask2 ^=t[i]
    }
    <- c
    return bitMask1 ^ bitMask2
}
