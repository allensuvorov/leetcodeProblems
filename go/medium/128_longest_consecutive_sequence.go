func longestConsecutive(nums []int) int {
    longest := 0
    numSet := make(map[int]bool,len(nums))
    for _, n := range nums {
        numSet[n] = true
    }
    
    for _, n := range nums {
        if !numSet[n - 1] { // is it the start of a sequence
            length := 1
            delete(numSet, n)
            for numSet[n + length]{
                delete(numSet, n + length)
                length++
            }
            if longest < length {
                longest = length
            }
        }
    }
    return longest
}
