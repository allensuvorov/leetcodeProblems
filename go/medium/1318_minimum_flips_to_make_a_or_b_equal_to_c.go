func minFlips(a int, b int, c int) int {
    or := a | b
    xor := or ^ c // all flips, single and double
    and := a & b & xor // double flips that match all flips
    return bits.OnesCount(uint(xor)) + bits.OnesCount(uint(and)) 
}
