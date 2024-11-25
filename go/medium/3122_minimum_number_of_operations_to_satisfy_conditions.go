/*
Medium - 3122. Minimum Number of Operations to Satisfy Conditions

Time: Big O (n)
Space: Big O (1)

Approach: Let's imagine this is a story about a lazy gardner.

* There are n gardens on a street. Each garden has m different flowers of garden variaty (10).
* Each owner asked: I want all my m flowers to be the same, but different from my neighbors.
* What's the minimum total number of flowers the gardner needs to change?

Solution: We are trying to keep each neighbor happy, whily being lazy. 
* Let's count cost of each solution in the current garden and add best option from previous garden.
* That will show best solution for each flower type so far.
*/

func minimumOperations(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])

    const digits = 10 // number of possible digits

    prev := [digits]int{} // prev best solution for each number

    for c := range n {
        curr := [digits]int{}
        for r := range m { // curr count 
            curr[grid[r][c]]++
        }
        for i := range digits { // curr count + best in prev for any other number
            maxPrev := 0
            for j := range digits {
                if i != j {
                    maxPrev = max(maxPrev, prev[j])
                }
            }
            curr[i] += maxPrev
        }
        prev = curr
    }

    maxPrev := 0
    for i := range digits {
        maxPrev = max(maxPrev, prev[i])
    }
    return m * n - maxPrev
}
