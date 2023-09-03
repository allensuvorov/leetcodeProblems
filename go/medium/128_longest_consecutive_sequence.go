func longestConsecutive(nums []int) int {
    longest := 0
    numSet := map[int]bool{}
    for _, n := range nums {
        numSet[n] = true
    }
    
    for _, n := range nums {
        if !numSet[n - 1] { // is it the start of a sequence
            length := 1
            for numSet[n + length]{
                length++
            }
            if longest < length {
                longest = length
            }
        }
    }
    return longest
}
