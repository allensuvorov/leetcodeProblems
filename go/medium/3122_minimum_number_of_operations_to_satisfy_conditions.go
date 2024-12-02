/*
Medium - 3122. Minimum Number of Operations to Satisfy Conditions

Time: Big O (n)
Space: Big O (1)

Approach: Let's imagine this is a story about a lazy gardner.

* There are n gardens on a street. Each garden has m different flowers of garden variaty (10).
* Each owner asked: I want all my m flowers to be the same, but different from my neighbors.
* What's the minimum total number of flowers the gardner needs to change?

Solution: We are trying to keep each neighbor happy, whily being lazy. 
* Let's count cost (or actually savings) of each solution in the current garden 
* and add best available option from previous garden.
* That will show best solution for each flower type for current garden so far.
*/

func minimumOperations(grid [][]int) int {
    rows, cols := len(grid), len(grid[0])
    const digitsVariety = 10 // number of possible digits 0-9

    prev := make([]int, digitsVariety) // prev best solution for each number

    for c := range cols {
        curr := make([]int, digitsVariety)
        for r := range rows { // curr counts 
            curr[grid[r][c]]++
        }
        for digit := range curr { // curr count + best in prev for any other number
            maxPrev := 0
            for j := range prev {
                if digit != j {
                    maxPrev = max(maxPrev, prev[j])
                }
            }
            curr[digit] += maxPrev
        }
        copy(prev, curr)
    }

    maxPrev := 0
    for _, v := range prev {
        maxPrev = max(maxPrev, v)
    }
    return rows * cols - maxPrev
}
