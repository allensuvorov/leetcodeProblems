package easy

// my solution
func countOdds(low int, high int) int {
    rangeLength := high - low + 1

    if rangeLength % 2 == 0 || low % 2 == 0{
        return rangeLength / 2
    }

    return rangeLength / 2 + 1
}

// some smarty solution
func countOdds(low int, high int) int {
    return (high + 1) / 2 - low / 2
}
